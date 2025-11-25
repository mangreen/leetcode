package q1182_ShortestDistanceToTargetColor

import (
	"math"
	"sort"
)

func shortestDistanceColor(colors []int, queries [][]int) []int {
	// 建立顏色到位置的映射
	colorMap := make(map[int][]int)
	for i, c := range colors {
		colorMap[c] = append(colorMap[c], i)
	}

	// 處理每個查詢
	ans := make([]int, len(queries))
	for i, q := range queries {
		idx, target := q[0], q[1]
		if pos, ok := colorMap[target]; ok {
			// 二分搜尋找到最接近 idx 的位置
			j := sort.SearchInts(pos, idx)
			minDist := math.MaxInt32
			// 檢查左邊位置
			if j > 0 {
				minDist = idx - pos[j-1]
			}
			// 檢查右邊位置
			if j < len(pos) {
				minDist = min(minDist, pos[j]-idx)
			}
			ans[i] = minDist
		} else {
			ans[i] = -1
		}
	}
	return ans
}

/*
ex.
q=[2, 1]

m={1: [0, 1, 3], 2: [2], 3: [4]}
             ↳j=sort.SearchInts(m[1], 2)=2
		 idx=2
		 ↓
c=[1, 1, 2, 1, 3]
      ↑     ↑ 
pos[j-1]=1  pos[j]=3
minDist=min(2-1, 3-2)=1
*/
