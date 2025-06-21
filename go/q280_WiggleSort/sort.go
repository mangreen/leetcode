package q280_WiggleSort

import "sort"

func wiggleSort2(nums []int) []int {
	if len(nums) <= 2 {
		return []int{}
	}

	sort.Ints(nums)

	for i := 2; i < len(nums); i += 2 {
		nums[i-1], nums[i] = nums[i], nums[i-1]
	}

	return nums
}
/*
[3, 5, 2, 1, 6, 4]

sorted
       i
[1, 2, 3, 4, 5, 6]

             i
[1, 3, 2, 4, 5, 6]

                   i
[1, 3, 2, 5, 4, 6]
*/
