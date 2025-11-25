package q256_PaintHouse

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_minCost_1(t *testing.T) {
	assert := assert.New(t)

	
    input := [][]int{{17,2,17},{16,16,5},{14,3,19}}
	ans := minCost(input)

	if assert.NotNil(ans) {
		assert.Equal(10, ans, "[[17,2,17],[16,16,5],[14,3,19]] MinCost: 2 + 5 + 3 = 10")
	}
}