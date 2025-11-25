package q261_GraphValidTree

import "slices"

func validTree2(n int, edges [][]int) bool {
	// Go 1.23+ 新增 lices.Repeat
	roots := slices.Repeat([]int{-1}, n)

	for _, e := range edges {
		x := find(roots, e[0])
		y := find(roots, e[1])

		if x == y {
			return false
		}

		roots[x] = y
	}

	// 樹的邊數量必定比節點數量少1
	return len(edges) == n-1
}

func find(roots []int, cur int) int {
	for roots[cur] != -1 {
		cur = find(roots, roots[cur])
	}

	return cur
}
