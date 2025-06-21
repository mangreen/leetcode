package q1644_LowestCommonAncestorOfABinaryTreeII

import (
	"fmt"
)

type TreeNode struct {
	Val int
 	Left *TreeNode
 	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	cnt := 0
	res := DFS(root, p, q, &cnt)

	if cnt == 2 {
		fmt.Println(res.Val)
		return res
	}

	return nil
}

func DFS(n, p, q *TreeNode, cnt *int) *TreeNode {
	if n == nil {
		return nil
	}

	l := DFS(n.Left, p, q, cnt)
	r := DFS(n.Right, p, q, cnt)

	if n == p || n == q {
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