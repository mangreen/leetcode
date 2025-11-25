package q256_PaintHouse

import "math"

func minCost(costs [][]int) int {
	if len(costs) == 0 || len(costs[0]) == 0 {
		return 0
	}

	min1, min2 := 0, 0
	idx := -1

	for i, row := range costs {
		m1, m2 := math.MaxInt, math.MaxInt
		jdx := -1

		for j := range row {
            // 找出最小和次小的，最小的要记录下标，方便下一轮判断
			MN := min1
			if j == idx {
				MN = min2
			}

			c := costs[i][j] + MN

			if c < m1 {
				m2 = m1
				m1 = c
				jdx = j
			} else if c < m2 {
				m2 = c
			}
		}

        min1 = m1
        min2 = m2
        idx = jdx
	}

    return min1
}

/*
ex1.
costs=[[1,5,3],[2,9,4]]
min1=0
min2=0
idx=-1

i=0
        j
costs=[[1,5,3],[2,9,4]] 
c=1+0=1 < m1=MaxInt
m1=MaxInt ⭢ 1
m2=MaxInt ⭢ MaxInt
jdx=-1 ⭢ 0

          j
costs=[[1,5,3],[2,9,4]]
c=5+0=5 > m1=1
        < m2=MaxInt
m1=1
m2=MaxInt ⭢ 5
jdx=0

            j
costs=[[1,5,3],[2,9,4]]
c=3+0=3 > m1=1
        < m2=5
m1=1
m2=5 ⭢ 3
jdx=0


min1=m1=1
min2=m2=3
idx=0


i=1
                j==jdx=0
costs=[[1,5,3],[2,9,4]] 
c=2+3=5 < m1=MaxInt
m1=MaxInt ⭢ 5
m2=MaxInt ⭢ MaxInt
jdx=-1 ⭢ 0

                  j
costs=[[1,5,3],[2,9,4]]
c=9+1=10 > m1=5
         < m2=MaxInt
m1=5
m2=MaxInt ⭢ 10
jdx=0

                    j
costs=[[1,5,3],[2,9,4]]
c=4+1=5 == m1=5
        < m2=10
m1=5
m2=10 ⭢ 5
jdx=0


min1=m1=5 ⭢ ans
min2=m2=5
idx=0
*/
