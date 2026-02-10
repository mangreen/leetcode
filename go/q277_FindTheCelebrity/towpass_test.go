package q277_FindTheCelebrity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_findCelebrity_1(t *testing.T) {
	assert := assert.New(t)

	/*
	  0   3 
      ⭣ ↘ ⭣ 
	  1 ⭢ 2
	*/
	graph = [][]bool{
		{false, true, true, false},
		{false, false, true, false},
		{false, false, false, false},
		{false, false, true, false},
	}
	ans := findCelebrity(4)

	if assert.NotNil(ans) {
		assert.Equal(2, ans, "The celebrity is person 2")
	}
}

func Test_findCelebrity_2(t *testing.T) {
	assert := assert.New(t)

	/*
		0
       ↗ 
	  1 ⭠ 2
	*/
	graph = [][]bool{
		{false, false, false},
		{true, false, false},
		{false, true, false},
	}
	ans := findCelebrity(3)

	if assert.NotNil(ans) {
		assert.Equal(-1, ans, "There is no celebrity")
	}
}