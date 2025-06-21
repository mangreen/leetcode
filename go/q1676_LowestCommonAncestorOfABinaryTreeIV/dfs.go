package q1676_LowestCommonAncestorOfABinaryTreeIV

import (
	"slices"
)

type TreeNode struct {
	Val int
 	Left *TreeNode
 	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, nodes []*TreeNode) *TreeNode {
	s := []int{}
	for _, n := range nodes{
		s = append(s, n.Val)
	}

	cnt := 0
	res := DFS(root, s, &cnt)

	if cnt == len(s) {
		return res
	}

	return nil
}

func DFS(n *TreeNode, s []int, cnt *int) *TreeNode {
    if n == nil {
		return nil
	}

	l := DFS(n.Left, s, cnt)
	r := DFS(n.Right, s, cnt)

	if slices.Contains(s, n.Val) {
        *cnt += 1
        return n
    }

    if l!=nil && r!=nil {
        return n
    }

	if l != nil {
		return l
	}

	return r
}