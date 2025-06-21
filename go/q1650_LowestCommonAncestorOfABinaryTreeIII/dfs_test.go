package q1650_LowestCommonAncestorOfABinaryTreeIII

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestor_1(t *testing.T) {
	assert := assert.New(t)

	/*
      2
     / \
    1   4
       /
      3
    */
    
    n3 := &TreeNode{3, nil, nil, nil}
    
    n4 := &TreeNode{4, n3, nil, nil}
    n3.Parent = n4
    
    n1 := &TreeNode{1, nil, nil, nil}
    
    n2 := &TreeNode{2, n1, n4, nil}
    n4.Parent = n2
    n1.Parent = n2

	ans := lowestCommonAncestor(n1, n3)

	assert.Equal(ans.Val, 2, "Node 2 are the LCA of p=1 & q=3")
}

func Test_lowestCommonAncestor_2(t *testing.T) {
	assert := assert.New(t)

	/*
      2
     / \
    1   4
       /
      3
    */
    
    n3 := &TreeNode{3, nil, nil, nil}
    
    n4 := &TreeNode{4, n3, nil, nil}
    n3.Parent = n4
    
    n1 := &TreeNode{1, nil, nil, nil}
    
    n2 := &TreeNode{2, n1, n4, nil}
    n4.Parent = n2
    n1.Parent = n2

	ans := lowestCommonAncestor(n4, n3)

	assert.Equal(ans.Val, 4, "Node 4 are the LCA of p=4 & q=3")
}