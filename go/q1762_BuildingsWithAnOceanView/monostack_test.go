package q1762_BuildingsWithAnOceanView

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findBuildings_1(t *testing.T) {
	assert := assert.New(t)

	input := []int{4, 2, 3, 1}

	ans := findBuildings(input)

	assert.Equal([]int{0, 2, 3}, ans, "Building 1 (0-indexed) does not have an ocean view because building 2 is taller.")
}

func Test_findBuildings_2(t *testing.T) {
	assert := assert.New(t)

	input := []int{4, 3, 2, 1}

	ans := findBuildings(input)

	assert.Equal([]int{0, 1, 2, 3}, ans, "All the buildings have an ocean view.")
}

func Test_findBuildings_3(t *testing.T) {
	assert := assert.New(t)

	input := []int{1, 3, 2, 4}

	ans := findBuildings(input)

	assert.Equal([]int{3}, ans, "Only building 3 has an ocean view.")
}
