package q1762_BuildingsWithAnOceanView

import ()

/*
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
*/

func findDistance2(root *TreeNode, p int, q int) int {
	lca := findLCA(root, p, q)

	distP := findLevel(lca, p)
	distQ := findLevel(lca, q)

	return distP + distQ
}

func findLCA(n *TreeNode, p, q int) *TreeNode {
	if n == nil || n.Val == p || n.Val == q {
		return n
	}

	left := findLCA(n.Left, p, q)
	right := findLCA(n.Right, p, q)

	if left != nil && right != nil {
		return n
	}

	if left != nil {
		return left
	}

	return right
}

func findLevel(n *TreeNode, target int) int {
	if n == nil {
		return -1
	}

	if n.Val == target {
		return 0
	}

	left := findLevel(n.Left, target)
	right := findLevel(n.Right, target)
	if left == -1 && right == -1 {
		return -1
	}	

	return 1 + max(left, right)
}