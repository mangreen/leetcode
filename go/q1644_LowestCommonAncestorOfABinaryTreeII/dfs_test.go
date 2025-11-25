package q1644_LowestCommonAncestorOfABinaryTreeII

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

	p := n6
	q := n1

	ans := lowestCommonAncestor(n3, p, q)

	assert.Equal(ans.Val, 3, "Node 3 are the LCA of p=6 & q=1")
}

func Test_lowestCommonAncestor_2(t *testing.T) {
	assert := assert.New(t)

	n6 := &TreeNode{6, nil, nil}
	n5 := &TreeNode{5, n6, nil}
	n1 := &TreeNode{1, nil, nil}
	n3 := &TreeNode{3, n5, n1}

	p := n5
	q := &TreeNode{10, nil, nil}

	ans := lowestCommonAncestor(n3, p, q)

	assert.Nil(ans, nil, "Should be nil, q=10 doesn't exist")
}