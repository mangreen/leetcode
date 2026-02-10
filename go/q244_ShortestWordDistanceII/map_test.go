package q244_ShortestWordDistanceII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_map_ex1(t *testing.T) {
	assert := assert.New(t)

	wordsDict := []string{"a", "b", "c", "a", "d", "b"}

	wd := Constructor(wordsDict)

	ans := wd.Shortest("a", "b")
	assert.Equal(1, ans, "The shortest distance between 'a' and 'b' is 1.")

	ans = wd.Shortest("c", "d")
	assert.Equal(2, ans, "The shortest distance between 'c' and 'd' is 2.")
}
