package q256_PaintHouse

func minCost(costs [][]int) int {
	if len(costs) == 0 || len(costs[0]) == 0 {
		return 0
	}

	r, g, b := costs[0][0], costs[0][1], costs[0][2]
	for _, c := range costs[1:] {
		rNew := c[0] + min(g, b)
		gNew := c[1] + min(r, b)
		bNew := c[2] + min(r, g)
		r, g, b = rNew, gNew, bNew
	}

	return min(r, min(g, b))
}
/*
Time: O(n)
Space: O(1)

ex1.
        r g b  
costs=[[1,5,3],[2,9,4]]
i=1
rNew=2+min(5,3)=5
gNew=9+min(1,3)=10
bNew=4+min(1,5)=5

return min(5,10,5)=5

ex2.
		r  g  b   
costs=[[17,2,17],[16,16,5],[14,3,19]]

i=1
rNew=16+min(2,17)=18
gNew=16+min(17,17)=33
bNew=5+min(17,2)=7

i=2
rNew=14+min(33,7)=21
gNew=3+min(18,7)=10
bNew=19+min(18,33)=37

return min(21,10,37)=10
*/
