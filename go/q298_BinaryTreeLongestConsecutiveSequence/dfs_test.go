package q298_BinaryTreeLongestConsecutiveSequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dfs1_ex1(t *testing.T) {
	assert := assert.New(t)

	var n5 = &TreeNode{Val: 5}
	var n4 = &TreeNode{Val: 4, Right: n5}
	var n2 = &TreeNode{Val: 2}
	var n3 = &TreeNode{Val: 3, Left: n2, Right: n4}
	var n1 = &TreeNode{Val: 1, Right: n3}

	ans := longestConsecutive(n1)

	/*
		1
		 \
		  3 
		 / \
		2   4
			 \
			  5
	*/
	assert.Equal(3, ans, "Longest consecutive sequence path is 3-4-5, so return 3.")
}

func Test_dfs1_ex2(t *testing.T) {
	assert := assert.New(t)

	var n1 = &TreeNode{Val: 1}
	var n2 = &TreeNode{Val: 2, Left: n1}
	var n3 = &TreeNode{Val: 3, Left: n2}
	var n0 = &TreeNode{Val: 2, Right: n3}

	ans := longestConsecutive(n0)

	/*
		2
		 \
		  3 
		 / 
		2   
	   /		
	  1		
	*/
	assert.Equal(2, ans, "Longest consecutive sequence path is 2-3, not 3-2-1, so return 2.")
}
