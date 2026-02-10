package q360_SortTransformedArray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twopoint_ex1(t *testing.T) {
	assert := assert.New(t)

	ans := sortTransformedArray([]int{-4, -2, 2, 4}, 1, 3, 5)
	
	expect := []int{3, 9, 15, 33}

	assert.Equal(expect, ans, "[-4, -2, 2, 4] a=1 b=3 c=5 ⭢ [3, 9, 15, 33]")
}

func Test_twopoint_ex2(t *testing.T) {
	assert := assert.New(t)

	ans := sortTransformedArray([]int{-4,-2,2,4}, -1, 3, 5)
	
	expect := []int{-23, -5, 1, 7}

	assert.Equal(expect, ans, "[-4, -2, 2, 4] a=-1 b=3 c=5 ⭢ [-23, -5, 1, 7]")
}
