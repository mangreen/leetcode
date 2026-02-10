package q1762_BuildingsWithAnOceanView

import (
	"slices"
)

func maxLength(ribbons []int, k int) int {
	L, R := 1, slices.Max(ribbons)

	ans := 0

	for L <= R {
		mid := (L + R) / 2

		if canCut(ribbons, k, mid) {
			ans = mid
			L = mid + 1
		} else {
			R = mid - 1
		}
	}

	return ans
}

func canCut(ribbons []int, k, curL int) bool {
	cnt := 0

	for _, r := range ribbons {
		cnt += r/curL

		if cnt >= k {
			return true
		}
	}

	return false
}
