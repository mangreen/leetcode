package q625_MinimumFactorization

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_greedy_68(t *testing.T) {
	assert := assert.New(t)

	input := 48
	ans := smallestFactorization(input)
	assert.Equal(68, ans, "68 is the answer because 6*8 = 48")
}

func Test_greedy_35(t *testing.T) {
	assert := assert.New(t)

	input := 15
	ans := smallestFactorization(input)
	assert.Equal(35, ans, "35 is the answer because 3*5 = 15")
}