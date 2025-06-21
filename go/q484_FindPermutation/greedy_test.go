package q484_FindPermutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findPermutation_inputI(t *testing.T) {
	assert := assert.New(t)

	input := "I"
	ans := findPermutation(input)

	assert.ElementsMatch([]int{1,2}, ans, "[1,2] is the only legal permutation.")
}

func Test_findPermutation_inputDI(t *testing.T) {
	assert := assert.New(t)

	input := "DI"
	ans := findPermutation(input)

	assert.ElementsMatch([]int{2,1,3}, ans, "Both [2,1,3] and [3,1,2] can construct the secret signature 'DI'.")
}
