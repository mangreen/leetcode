package q1182_ShortestDistanceToTargetColor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_binSearchAndMap_1(t *testing.T) {
	assert := assert.New(t)

	colors := []int{1,1,2,1,3,2,2,3,3} 
	queries := [][]int{{1,3},{2,2},{6,1}}
	ans := shortestDistanceColor(colors, queries)

	assert.Equal([]int{3,0,3}, ans, "Output should be [3,0,3]")
}

func Test_binSearchAndMap_2(t *testing.T) {
	assert := assert.New(t)

	colors := []int{1,2} 
	queries := [][]int{{0,3}}
	ans := shortestDistanceColor(colors, queries)

	assert.Equal([]int{-1}, ans, "There is no 3 in the array.")
}
