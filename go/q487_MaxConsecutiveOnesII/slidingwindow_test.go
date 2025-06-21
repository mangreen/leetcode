package q487_MaxConsecutiveOnesII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMaxConsecutiveOnes_output4(t *testing.T) {
	assert := assert.New(t)

	input := []int{ 1, 0, 1, 1, 0 }
	ans := findMaxConsecutiveOnes(input)

	assert.Equal(4, ans, "Flip 1st 0 to get the max consecutive 1s is 4.")
}

func Test_findMaxConsecutiveOnes_output3(t *testing.T) {
	assert := assert.New(t)

	input := []int{ 0, 0, 1, 1, 0 }
	ans := findMaxConsecutiveOnes(input)

	assert.Equal(3, ans, "Flip 2nd 0 to get the max consecutive 1s is 3.")
}

func Test_findMaxConsecutiveOnes_output5(t *testing.T) {
	assert := assert.New(t)

	input := []int{ 1, 1, 1, 1, 0 }
	ans := findMaxConsecutiveOnes(input)

	assert.Equal(5, ans, "Flip 1st/last 0 to get the max consecutive 1s is 5.")
}
