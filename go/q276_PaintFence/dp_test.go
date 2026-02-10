package q276_PaintFence

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_dp_ex1(t *testing.T) {
	assert := assert.New(t)

	result := numWays(3, 2)
	assert.Equal(6, result, "Expected 6 ways for n=3, k=2")
}

func Test_dp_ex2(t *testing.T) {
	assert := assert.New(t)

	result := numWays(1, 1)
	assert.Equal(1, result, "Expected 1 way for n=1, k=1")
}