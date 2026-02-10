package q270_ClosestBinarySearchTreeValue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_binsearch_ex1(t *testing.T) {
	assert := assert.New(t)

	input := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,	
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}

	ans := closestValue(input, 3.714286)

	assert.Equal(4, ans, "Expected value should be 4.")
}

func Test_binsearch_ex2(t *testing.T) {
	assert := assert.New(t)

	input := &TreeNode{
		Val: 1,
	}

	ans := closestValue(input, 1.0) 
	
	assert.Equal(1, ans, "Expected value should be 1.")
}
