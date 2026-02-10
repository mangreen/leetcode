package q270_ClosestBinarySearchTreeValue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dfs_ex1(t *testing.T) {
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

	ans := closestKValues(input, 3.714286, 2)

	assert.Equal([]int{3, 4}, ans, "Expected values should be [3, 4].")
}

func Test_dfs_ex2(t *testing.T) {
	assert := assert.New(t)

	input := &TreeNode{
		Val: 1,
	}

	ans := closestKValues(input, 1.0, 1) 
	
	assert.Equal([]int{1}, ans, "Expected value should be [1].")
}
