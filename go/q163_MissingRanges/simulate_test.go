package q163_MissingRanges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_simulate_ex1(t *testing.T) {
	assert := assert.New(t)

	input := []int{0, 1, 3, 50, 75}
	lower := 0
	upper := 99
	ans := findMissingRanges(input, lower, upper)

	assert.Equal([][]int{{2, 2}, {4, 49}, {51, 74}, {76, 99}}, ans, "ab can OneEditDistance to be acb")
}

func Test_simulate_ex2(t *testing.T) {
	assert := assert.New(t)

	input := []int{-1}
	lower := -1
	upper := -1
	ans := findMissingRanges(input, lower, upper)
	
	assert.Equal([][]int(nil), ans, "There are no missing ranges since there are no missing numbers.")
}