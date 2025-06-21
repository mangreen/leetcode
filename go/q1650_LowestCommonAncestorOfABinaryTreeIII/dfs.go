package q1650_LowestCommonAncestorOfABinaryTreeIII

import (
	"fmt"
)

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

func lowestCommonAncestor(p, q *TreeNode) *TreeNode {
	res := DFS(p, q)

	if res != nil {
		fmt.Println(res.Val)
		return res
	}

	return nil
}

func DFS(p, q *TreeNode) *TreeNode {
	a, b := p, q
	for a != b {
		if a.Parent != nil {
			a = a.Parent
		} else {
			a = q
		}
		
		if b.Parent != nil {
			b = b.Parent
		} else {
			b = p
		}
	}
	return a
}
