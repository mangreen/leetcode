package q250_CountUnivalueSubtrees


import (
	"testing"

	"github.com/stretchr/testify/assert"
)



func Test_dfs_ex1(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{Val: 5}},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{Val: 5},
		},
	}
	/*
			   5
			 /   \
			1	  5
		   /       \
		  5	   		5
	*/
	
	assert.Equal(4, countUnivalSubtrees(root), "should be 4")
}

func Test_dfs_ex2(t *testing.T) {
	assert := assert.New(t)

	var root *TreeNode = nil
	
	assert.Equal(0, countUnivalSubtrees(root), "should be 0")
}

func Test_dfs_ex3(t *testing.T) {
	assert := assert.New(t)
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{Val: 5},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{Val: 5},
		},
	}
	/*
			   5
			 /   \
			5	  5
		   / \     \ 
		  5	  5 	5	
	*/

	assert.Equal(6, countUnivalSubtrees(root), "should be 6")
}
