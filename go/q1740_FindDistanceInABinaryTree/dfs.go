package q1762_BuildingsWithAnOceanView

import ()

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findDistance(root *TreeNode, p int, q int) int {
	var p2P, p2Q string

	DFS(root, p, q, &p2P, &p2Q, nil)

	i := 0
	for i < min(len(p2P), len(p2Q)) && p2P[i] == p2Q[i] {
		i++
	}

	return len(p2P[i:]) + len(p2Q[i:])
}

func DFS(n *TreeNode, p, q int, p2P, p2Q *string, path []byte) {
	if n == nil {
		return
	}

	if n.Val == p {
		*p2P = string(path)
	}

	if n.Val == q {
		*p2Q = string(path)
	}

	
	DFS(n.Left, p, q, p2P, p2Q, append(path, 'L'))
	DFS(n.Right, p, q, p2P, p2Q, append(path, 'R'))
}