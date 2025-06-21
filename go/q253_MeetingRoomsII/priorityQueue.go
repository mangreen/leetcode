package q252_MeetingRooms

import (
	"container/heap"
	"sort"
)

func minMeetingRooms2(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	h := &MinHeap{}
	heap.Init(h)

	for _, v := range intervals {
		if h.Len() == 0 {
			heap.Push(h, Pair{
				S: v[0],
				E: v[1],
			})
			
			continue
		} 

		top := (*h)[0]

		if top.E <= v[0] {
			heap.Pop(h)
		}

		heap.Push(h, Pair{
			S: v[0],
			E: v[1],
		})
	}

	return h.Len()
}

type Pair struct {
	S int
	E int
}

type MinHeap []Pair

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].E < h[j].E
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() any {
	old := *h
	N := len(old)
	x := old[N-1]
	*h = old[:N-1]
	return x
}
