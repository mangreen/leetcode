package q1762_BuildingsWithAnOceanView

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_binarysearch_1(t *testing.T) {
	assert := assert.New(t)

	input := []int{9, 7, 5}
	k := 3

	ans := maxLength(input, k)

	assert.Equal(5, ans, "Building 1 (0-indexed) does not have an ocean view because building 2 is taller.")
}

func Test_binarysearch_2(t *testing.T) {
	assert := assert.New(t)

	input := []int{7, 5, 9}
	k := 4

	ans := maxLength(input, k)

	assert.Equal(5, ans, "All the buildings have an ocean view.")
}

func Test_binarysearch_3(t *testing.T) {
	assert := assert.New(t)

	input := []int{5, 7, 9}
	k := 22

	ans := maxLength(input, k)

	assert.Equal(0, ans, "Only building 3 has an ocean view.")
}
