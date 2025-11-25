package main

// import (
// 	"fmt"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"
// )

// func main() {
//     fmt.Println("Starting traffic light simulation. Press Ctrl+C to stop.")

//     // 建立 context-like quit channel
//     quit := make(chan struct{})
//     sigCh := make(chan os.Signal, 1)
//     signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

//     // traffic light goroutine
//     go func() {
//         for {
//             select {
//                 case <-quit:
//                     fmt.Println("Traffic light simulation stopped.")
//                     return
                
//                 default:
//                     // Green light phase
//                     if !simulate(quit, Green, GreenDuration) {
//                         return
//                     }
                
//                     // Orange light phase
//                     if !simulate(quit, Orange, OrangeDuration) {
//                         return
//                     }
                
//                     // Red light phase
//                     if !simulate(quit, Red, RedDuration) {
//                         return
//                     }
//             }
//         }
//     }()

//     <-sigCh // 阻塞主 goroutine，等待訊號
//     fmt.Println("Received interrupt. Shutting down...")
//     close(quit) // 通知 goroutine 停止

//     // 等待一點時間讓 goroutine 結束
//     time.Sleep(1 * time.Second)
    
//     fmt.Println("os exit")
// 	os.Exit(0) // 正常退出
// }

// const (
//     GreenDuration  = 5 * time.Second
//     OrangeDuration = 2 * time.Second
//     RedDuration    = 6 * time.Second
// )

// const (
//     Green  = "Green"
//     Orange = "Orange"
//     Red    = "Red"
// )

// // simulate simulates a single light phase and prints the countdown.
// // return false 表示被中斷
// func simulate(quit chan struct{}, color string, duration time.Duration) bool {
//     ticker := time.NewTicker(1 * time.Second) // 倒數計時器，每秒觸發一次
//     defer ticker.Stop()

//     timer := time.NewTimer(duration) // 剩餘時間計時器，到時結束該 color
//     defer timer.Stop()

//     for {
//         select {
//             case <-quit:
//                 fmt.Printf("Simulate interrupted during %s light!\n", color)
//                 return false
            
//             case <-ticker.C:
//                 if duration > 0 {
//                     fmt.Printf("Light %s remaining: %v\n", color, duration)
//                 }
//                 duration -= 1 * time.Second
//                 // 不回傳，繼續 for 迴圈
            
//             case <-timer.C:
//                 fmt.Printf("Light %s finished.\n", color)
//                 return true
//         }
//     }
// }