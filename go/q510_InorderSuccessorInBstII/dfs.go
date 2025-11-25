package q510_InorderSuccessorInBstII

import ()

type Node struct {
	Val int
 	Left *Node
 	Right *Node
    Parent *Node
}

func inorderSuccessor(node *Node) *Node {
    if node == nil {
        return nil
    }

    if node.Right != nil {
        node = node.Right

        for node.Left != nil {
            node = node.Left
        }

        return node
    } 
    
    for node != nil {
        if node.Parent == nil {
            return nil
        }

        if node == node.Parent.Left {
            return node.Parent
        }

        node = node.Parent
    }

    return node
}
/*
       5
      / \ 
    3     7
   / \   /
  2   4 6
 /
1

★ 如果 n.Right 存在, 找到 n.Right 最 Left 的節點
ex1.
n=3 ⭢ n.R=4 ⭢ n.L=nil ans=4
ex3.
n=5 ⭢ n.R=7 ⭢ n.L=6 ans=6

★ 否則往上找直到 n == n.Parent.Left
ex4.
n=1 == n.p2.L ⭢ ans=2

★ 不然就是 nil　
ex2.
n=7 ⭢ n.p=5 ⭢ n.p=nil ⭢ ans=nil
*/