package q163_MissingRanges

func findMissingRanges(nums []int, lower int, upper int) [][]int {
	var res [][]int
	
	next := lower
	
	for _, v := range nums {
		// v 大於 next，表示中間有缺漏區間
		if v > next {
			res = append(res, []int{next, v - 1})
		}

		// 更新 next 為 v 的下一個數字
		next = v + 1
	}

	// 最後檢查是否還有剩餘的缺漏區間 (還沒到 upper)
	if next <= upper {
		res = append(res, []int{ next, upper })
	}

	return res
}
/*
ex.

[0, 1, 5], lower=0, upper=9
next=0
res=[]

v=0 ≥ next=0 ⭢ next=1	

v=1 ≥ next=1 ⭢ next=2

v=5 ≥ next=2 ⭢ res=[[2,4]], next=6

結束迴圈

next=6 ≤ upper=9 ⭢ res=[[2,4],[6,9]]
*/