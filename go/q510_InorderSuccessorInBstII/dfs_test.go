package q510_InorderSuccessorInBstII

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_inorderSuccessor_1(t *testing.T) {
	assert := assert.New(t)

	n2 := &Node{2, nil, nil, nil}
    n1 := &Node{1, nil, nil, n2}
    n3 := &Node{3, nil, nil, n2}
	n2.Left = n1
	n2.Right = n3

	ans := inorderSuccessor(n1)

	if assert.NotNil(ans) {
		assert.Equal(2, ans.Val, "Node 1's in-order successor node is 2")
	}
}

func Test_inorderSuccessor_2(t *testing.T) {
	assert := assert.New(t)

	n5 := &Node{5, nil, nil, nil}
	n3 := &Node{3, nil, nil, n5}
	n6 := &Node{6, nil, nil, n5}
	n5.Left = n3
	n5.Right = n6

	n2 := &Node{2, nil, nil, n3}
	n4 := &Node{4, nil, nil, n3}
	n3.Left = n2
	n3.Right = n4

    n1 := &Node{1, nil, nil, n2}
	n2.Left = n1

	ans := inorderSuccessor(n6)

	assert.Nil(ans, "There is no in-order successor of the current node 6")
}

func Test_inorderSuccessor_3(t *testing.T) {
	assert := assert.New(t)

	n5 := &Node{5, nil, nil, nil}
	n3 := &Node{3, nil, nil, n5}
	n6 := &Node{6, nil, nil, n5}
	n5.Left = n3
	n5.Right = n6

	n2 := &Node{2, nil, nil, n3}
	n4 := &Node{4, nil, nil, n3}
	n3.Left = n2
	n3.Right = n4

    n1 := &Node{1, nil, nil, n2}
	n2.Left = n1

	ans := inorderSuccessor(n5)

	if assert.NotNil(ans) {
		assert.Equal(6, ans.Val, "Node 5's in-order successor node is 6")
	}
}
