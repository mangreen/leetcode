package q156_BinaryTreeUpsideDown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_iterate_ex1(t *testing.T) {
	assert := assert.New(t)

	input := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,	
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	ans := upsideDownBinaryTree(input)

	expected := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 5,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}

	assert.Equal(expected, ans, "Expected tree structure should be match.")
}

func Test_iterate_ex2(t *testing.T) {
	assert := assert.New(t)

	input := &TreeNode{ Val: 1 }

	ans := upsideDownBinaryTree(input) 
	
	expected := input

	assert.Equal(expected, ans, "Expected input tree to be returned for single node tree.")
}
