package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Starting traffic light simulation. Press Ctrl+C to stop.")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	wg := new(sync.WaitGroup)
	for c, d := range Colors {
		wg.Add(1)

		go func(w *sync.WaitGroup, color string, duration time.Duration) {
			defer w.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Traffic light simulation stopped.")
					return

				default:
					if !simulate(ctx, c, d) {
						return
					}
				}
			}
		}(wg, c, d)
	}

	// <-ctx.Done() // 阻塞主 goroutine，等待訊號
	wg.Wait()
	fmt.Println("Received interrupt. Shutting down...")
	cancel() // 取消上下文，通知其他 goroutine 停止工作

	time.Sleep(1 * time.Second) // 等待其他 goroutine 完成清理工作

	fmt.Println("os exit")
	os.Exit(0) // 正常退出
}

const (
	GREEN_DURATION  = 5 * time.Second
	ORANGE_DURATION = 2 * time.Second
	RED_DURATION    = 6 * time.Second
)

const (
	GREEN  = "Green"
	ORANGE = "Orange"
	RED    = "Red"
)

var Colors = map[string]time.Duration{
	GREEN:  GREEN_DURATION,
	ORANGE: ORANGE_DURATION,
	RED:    RED_DURATION,
}

// simulate simulates a single light phase and prints the countdown.
func simulate(ctx context.Context, color string, duration time.Duration) bool {
	ticker := time.NewTicker(1 * time.Second) // 倒數計時器，每秒觸發一次
	defer ticker.Stop()

	timer := time.NewTimer(duration) // 剩餘時間計時器，到時結束該 color
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Simulate interrupted during %s light!\n", color)
			return false

		case <-ticker.C:
			if duration > 0 {
				fmt.Printf("Light %s remaining: %v\n", color, duration)
			}
			duration -= 1 * time.Second
			// 不回傳，繼續 for 迴圈

		case <-timer.C:
			fmt.Printf("Light %s finished.\n", color)
			return true
		}
	}
}
