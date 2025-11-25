package q370_RangeAddition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getModifiedArray_1(t *testing.T) {
	assert := assert.New(t)

	length := 5
	updates := [][]int{{1,3,2},{2,4,3},{0,2,-2}}
	ans := getModifiedArray(length, updates)

	assert.Equal([]int{-2,0,3,5,3}, ans, "Output: [-2,0,3,5,3]")
}

func Test_getModifiedArray_2(t *testing.T) {
	assert := assert.New(t)

	length := 10
	updates := [][]int{{2,4,6},{5,6,8},{1,9,-4}}
	ans := getModifiedArray(length, updates)

	assert.Equal([]int{0,-4,2,2,2,4,4,-4,-4,-4}, ans, "Output: [0,-4,2,2,2,4,4,-4,-4,-4]")
}
