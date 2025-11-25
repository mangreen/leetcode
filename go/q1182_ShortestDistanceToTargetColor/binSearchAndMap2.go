package q1182_ShortestDistanceToTargetColor

func shortestDistanceColor2(colors []int, queries [][]int) []int {
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
			L, R := 0, len(pos)-1

			for L < R {
				mid := (L + R) / 2
				if pos[mid] == idx {
					ans[i] = 0
					break
				} else if pos[mid] < idx {
					L = mid
				} else {
					R = mid
				}
			}

			ans[i] = min(idx-pos[L], pos[R]-idx)
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
	L=1     R=3
minDist=min(2-1, 3-2)=1
*/
