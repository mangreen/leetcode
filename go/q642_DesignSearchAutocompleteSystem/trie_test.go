package q642_DesignSearchAutocompleteSystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_trie_1(t *testing.T) {
	assert := assert.New(t)

	sentences := []string{ "i love you", "island","ironman", "i love leetcode" }
    times := []int{ 5, 3, 2, 2 }
    sys := ConstructorTrie(sentences, times)
    
	ans := sys.Input('i')
	assert.Equal(3, len(ans), "There are 3 corresponding items")
	mockAns := []string{ "i love you", "island", "i love leetcode" }
	assert.Equal(mockAns, ans, "Ans should be [i love you, island, i love leetcode]")

	ans = sys.Input(' ')
	assert.Equal(2, len(ans), "There are 2 corresponding items")
	mockAns = []string{ "i love you", "i love leetcode" }
	assert.Equal(mockAns, ans, "Ans should be [i love you, i love leetcode]")
}

func Test_trie_2(t *testing.T) {
	assert := assert.New(t)

	sentences := []string{ "i love you", "island","ironman", "i love leetcode" }
    times := []int{ 5,3,2,2 }
    sys := ConstructorTrie(sentences, times)
    
	sys.Input('i')
	sys.Input(' ')
	ans := sys.Input('#')
	assert.Equal(0, len(ans), "There are 0 corresponding items")
}