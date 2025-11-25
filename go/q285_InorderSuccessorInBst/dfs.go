package q285_InorderSuccessorInBst

import ()

type TreeNode struct {
	Val int
 	Left *TreeNode
 	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    var ans *TreeNode
    
    for root != nil {
        if root.Val > p.Val {
            ans = root
            root = root.Left
        } else {
            root = root.Right
        }
    }

    return ans
}
/*
       5
      / \ 
    3     7
   / \   /
  2   4 6
 /
1

★ 如果 p < cur 假設 ans=cur 再往 Left 找
ex1. p=3 
c=r=5>p a=5 ⭢ r=3==p ⭢ r=4>p a=4 ⭢ r=nil
                               ↳ ans=4

★ 否則 cur <= p 往 右
ex3. p=5 
c=r=5==p ⭢ r=7>p a=6 ⭢ r=6==p ⭢ r=nil
                    ↳ ans=6

★ 不然就是 nil　
ex2. p=7 
c=r=7==p ⭢ r=nil
ans=nil
*/