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

	/*
        題目不保證 p 和 q 一定存在於樹中，
        所以需要先搜尋子樹來確認 p 和 q 是否都存在。
		如果只要找到其中一個就可以立即返回，
		有可能 p 或 q 存在其子樹而錯誤。
    */
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