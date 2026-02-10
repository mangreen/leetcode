package q1762_BuildingsWithAnOceanView

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dfs2_findDistance_1(t *testing.T) {
	assert := assert.New(t)

	n8	:= &TreeNode{Val: 8}
	n0 := &TreeNode{Val: 0}
	n1 := &TreeNode{Val: 1, Left: n0, Right: n8}	
	n7 := &TreeNode{Val: 7}
	n4 := &TreeNode{Val: 4}
	n2 := &TreeNode{Val: 2, Left: n7, Right: n4}
	n6 := &TreeNode{Val: 6}
	n5 := &TreeNode{Val: 5, Left: n6, Right: n2}
	root :=	&TreeNode{Val: 3, Left: n5, Right: n1} 

	ans := findDistance2(root, 5, 0)

	assert.Equal(3, ans, "There are 3 edges between 5 and 0: 5-3-1-0.")
}

func Test_dfs2_findDistance_2(t *testing.T) {
	assert := assert.New(t)

	n8	:= &TreeNode{Val: 8}
	n0 := &TreeNode{Val: 0}
	n1 := &TreeNode{Val: 1, Left: n0, Right: n8}	
	n7 := &TreeNode{Val: 7}
	n4 := &TreeNode{Val: 4}
	n2 := &TreeNode{Val: 2, Left: n7, Right: n4}
	n6 := &TreeNode{Val: 6}
	n5 := &TreeNode{Val: 5, Left: n6, Right: n2}
	root :=	&TreeNode{Val: 3, Left: n5, Right: n1} 

	ans := findDistance2(root, 5, 7)

	assert.Equal(2, ans, "There are 2 edges between 5 and 7: 5-2-7.")
}

func Test_dfs2_findDistance_3(t *testing.T) {
	assert := assert.New(t)

	n8	:= &TreeNode{Val: 8}
	n0 := &TreeNode{Val: 0}
	n1 := &TreeNode{Val: 1, Left: n0, Right: n8}	
	n7 := &TreeNode{Val: 7}
	n4 := &TreeNode{Val: 4}
	n2 := &TreeNode{Val: 2, Left: n7, Right: n4}
	n6 := &TreeNode{Val: 6}
	n5 := &TreeNode{Val: 5, Left: n6, Right: n2}
	root :=	&TreeNode{Val: 3, Left: n5, Right: n1} 

	ans := findDistance2(root, 5, 5)

	assert.Equal(0, ans, "The distance between a node and itself is 0.")
}
