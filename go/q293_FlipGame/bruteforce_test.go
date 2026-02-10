package q293_FlipGame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bruteforce_ex1(t *testing.T) {
	assert := assert.New(t)

	ans := generatePossibleNextMoves("++++")

	assert.Equal([]string{"--++", "+--+", "++--"}, ans, "++++  can flip to --++, +--+, ++--")
}

func Test_bruteforce_ex2(t *testing.T) {
	assert := assert.New(t)

	ans := generatePossibleNextMoves("+")

	assert.Equal([]string(nil), ans, "single + cannot be flipped")
}