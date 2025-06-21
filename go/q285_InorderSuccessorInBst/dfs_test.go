package q285_InorderSuccessorInBst

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_inorderSuccessor_1(t *testing.T) {
	assert := assert.New(t)

	
    n1 := &TreeNode{1, nil, nil}
    n3 := &TreeNode{3, nil, nil}
	n2 := &TreeNode{2, n1, n3}

	ans := inorderSuccessor(n2, n1)

	if assert.NotNil(ans) {
		assert.Equal(2, ans.Val, "Node 1's in-order successor node is 2")
	}
}

func Test_inorderSuccessor_2(t *testing.T) {
	assert := assert.New(t)

	n1 := &TreeNode{1, nil, nil}
	n2 := &TreeNode{2, n1, nil}
	n4 := &TreeNode{4, nil, nil}
	n3 := &TreeNode{3, n2, n4}
	n6 := &TreeNode{6, nil, nil}
	n5 := &TreeNode{5, n3, n6}

	ans := inorderSuccessor(n5, n6)

	assert.Nil(ans, "There is no in-order successor of the current node 6")
}

func Test_inorderSuccessor_3(t *testing.T) {
	assert := assert.New(t)

	n1 := &TreeNode{1, nil, nil}
	n2 := &TreeNode{2, n1, nil}
	n4 := &TreeNode{4, nil, nil}
	n3 := &TreeNode{3, n2, n4}
	n6 := &TreeNode{6, nil, nil}
	n5 := &TreeNode{5, n3, n6}

	ans := inorderSuccessor(n5, n5)

	if assert.NotNil(ans) {
		assert.Equal(6, ans.Val, "Node 5's in-order successor node is 6")
	}
}
