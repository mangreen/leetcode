package q294_FlipGameII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_map_ex1(t *testing.T) {
	assert := assert.New(t)

	ans := canWin("++++")

	assert.Equal(true, ans, "++++ can be flipped to --++ then opponent cannot win")
}

func Test_map_ex2(t *testing.T) {
	assert := assert.New(t)

	ans := canWin("+")

	assert.Equal(false, ans, "+ cannot be flipped")
}