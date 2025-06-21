package q1676_LowestCommonAncestorOfABinaryTreeIV

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestor_1(t *testing.T) {
	assert := assert.New(t)

	n6 := &TreeNode{6, nil, nil}
	n5 := &TreeNode{5, n6, nil}
	n1 := &TreeNode{1, nil, nil}
	n3 := &TreeNode{3, n5, n1}

	nodes := []*TreeNode{n1, n5, n6}

	ans := lowestCommonAncestor(n3, nodes)

	assert.Equal(ans.Val, 3, "Node 3 are the LCA of Node 5, 1, 6")
}

func Test_lowestCommonAncestor_2(t *testing.T) {
	assert := assert.New(t)

	n6 := &TreeNode{6, nil, nil}
	n5 := &TreeNode{5, n6, nil}
	n1 := &TreeNode{1, nil, nil}
	n3 := &TreeNode{3, n5, n1}

	n99 := &TreeNode{99, nil, nil}

	nodes := []*TreeNode{n1, n99}

	ans := lowestCommonAncestor(n3, nodes)

	assert.Nil(ans, nil, "Should be nil, Node 99 doesn't exist")
}
