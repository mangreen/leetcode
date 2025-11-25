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
		if h.Len() > 0 && (*h)[0].E <= v[0] {
			heap.Pop(h)
		}

		heap.Push(h, Pair{
			S: v[0],
			E: v[1],
		})
	}

	return h.Len()
}

/*
Time: O(NlogN) // 每次 heap.Push 要 O(logN), 操作 N 次
Space: O(N) // heap 存放 N 個會議

★ 為什麼需要 heap
[[1, 3], [2, 4], [3, 5]] 無法用單純的變數紀錄結束時間，因為 [1,3] 和 [2,4] 會同時存在

★ 為什麼需要用 end 排序
因為我們要找最早結束的會議，才能釋放出會議室給後續的會議使用
ex.
[[1,    4], 
   [2,3], 
     [3,  5]]
用 start 排序的話，會錯過 [2,3] 這個會議室可以給 [3,5] 使用的機會
*/

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
