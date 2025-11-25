package q1087_BraceExpansion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_backtrack_1(t *testing.T) {
	assert := assert.New(t)

	ans := expand("{a,b}c{d,e}f")

	assert.Equal([]string{"acdf","acef","bcdf","bcef"}, ans, "[acdf,acef,bcdf,bcef]")
}

func Test_backtrack_2(t *testing.T) {
	assert := assert.New(t)

	ans := expand("abcd")

	assert.Equal([]string{"abcd"}, ans, "[abcd]")
}
