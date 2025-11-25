package q1151_MinimumSwapsToGroupAll1sTogether

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_slidingwindows_1(t *testing.T) {
	assert := assert.New(t)

	input := []int{1,0,1,0,1} 
	ans := minSwaps(input)

	assert.Equal(1, ans, "The minimum is 1")
}

func Test_slidingwindows_2(t *testing.T) {
	assert := assert.New(t)

	input := []int{0,0,0,1,0}
	ans := minSwaps(input)

	assert.Equal(0, ans, "Since there is only one 1 in the array, no swaps are needed.")
}

func Test_slidingwindows_3(t *testing.T) {
	assert := assert.New(t)

	input := []int{1,0,1,0,1,0,0,1,1,0,1}
	ans := minSwaps(input)

	assert.Equal(3, ans, "3 swaps for [0,0,0,0,0,1,1,1,1,1,1]")
}