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
/*
為什麼需要 a=q & b=p 交換指向？

  r         r=a        r=b        r
 /         /          /          / 
p=a   ⭢   p=b   ⭢   p     ⭢   p=a=b
 \         \          \          \
  q=b       q          q=a        q  

補齊兩個節點到達根節點的距離差。
類似 142. Linked List Cycle II 找 cycle 起始點
*/

