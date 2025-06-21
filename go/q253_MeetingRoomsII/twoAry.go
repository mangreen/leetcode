package q252_MeetingRooms

import "sort"

func minMeetingRooms(intervals [][]int) int {
	starts := []int{}
	ends := []int{}

	for _, v := range intervals {
		starts = append(starts, v[0])
		ends = append(ends, v[1])
	}

	sort.Ints(starts)
	sort.Ints(ends)

	ans := 0
	endIdx := 0
	for i := range starts {
		if starts[i] <= ends[endIdx] {
			ans++
			continue
		}

		endIdx++
	}

	return ans
}
