package q362_DesignHitCounter

type HitCounter struct {
    hits []int
}

func Constructor() HitCounter {
    return HitCounter{
        hits: []int{},
    }
}

func (this *HitCounter) Hit(timestamp int) {
    this.hits = append(this.hits, timestamp)
}

func (this *HitCounter) GetHits(timestamp int) int {
    // Remove old hits from the front of the queue
    for len(this.hits) > 0 && this.hits[0] <= timestamp-300 {
        this.hits = this.hits[1:] // Effectively pop from the front
    }
    
    return len(this.hits)
}