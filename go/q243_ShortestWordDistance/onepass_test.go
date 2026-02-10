package q243_ShortestWordDistance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_onepass_ex1(t *testing.T) {
	assert := assert.New(t)

	wordsDict := []string{"practice", "makes", "perfect", "coding", "makes"}
	word1 := "coding"
	word2 := "practice"
	ans := shortestDistance(wordsDict, word1, word2)

	assert.Equal(3, ans, "The shortest distance between 'coding' and 'practice' is 3.")
}

func Test_onepass_ex2(t *testing.T) {
	assert := assert.New(t)

	wordsDict := []string{"practice", "makes", "perfect", "coding", "makes"}
	word1 := "makes"
	word2 := "coding"
	ans := shortestDistance(wordsDict, word1, word2)	

	assert.Equal(1, ans, "The shortest distance between 'makes' and 'coding' is 1.")
}