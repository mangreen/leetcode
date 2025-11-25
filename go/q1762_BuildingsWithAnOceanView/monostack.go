package q1762_BuildingsWithAnOceanView

import (
	
)

func findBuildings(heights []int) []int {
	stk := []int{ len(heights)-1 }

	for i:=len(heights)-2; i>=0; i-- {
		if heights[i] > heights[stk[0]] {
			stk = append([]int{i}, stk...)
		}
	}

	return stk
}