package q346_MovingAverageFromDataStream

type MovingAverage struct {
	windowSize int
	queue      []int // Using a slice as a queue
	sum        int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		windowSize: size,
		queue:      make([]int, 0, size), // Pre-allocate capacity for efficiency
		sum:        0,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	// If the queue is full, remove the oldest element
	if len(this.queue) == this.windowSize {
		this.sum -= this.queue[0]
		this.queue = this.queue[1:] // Remove the first element
	}

	// Add the new value to the queue and update the sum
	this.queue = append(this.queue, val)
	this.sum += val

	// Calculate and return the moving average
	return float64(this.sum) / float64(len(this.queue))
}
