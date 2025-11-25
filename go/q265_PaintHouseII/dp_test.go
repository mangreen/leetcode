package q256_PaintHouse

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_minCost_1(t *testing.T) {
	assert := assert.New(t)

	
    input := [][]int{{1,5,3},{2,9,4}}
	ans := minCost(input)

	if assert.NotNil(ans) {
		assert.Equal(5, ans, "[[1,5,3],[2,9,4]] MinCost: Minimum cost: 1 + 4 = 5 or 3 + 2 = 5")
	}
}