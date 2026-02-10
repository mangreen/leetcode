package q186_ReverseWordsInAStringII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_map_ex1(t *testing.T) {
	assert := assert.New(t)

	str := []byte("the sky is blue")
	reverseWords(str)
	
	assert.Equal([]byte("blue is sky the"), str, "'the sky is blue' are reversed to 'blue is sky the'")
}

func Test_map_ex2(t *testing.T) {
	assert := assert.New(t)

	str := []byte("  hello world  ")
	reverseWords(str)

	assert.Equal([]byte("  world hello  "), str, "'  hello world  ' are reversed to '  world hello  '")
}