package q246_StrobogrammaticNumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twopoint_ex1(t *testing.T) {
	assert := assert.New(t)

	ans := isStrobogrammatic("69")

	assert.Equal(true, ans, "The number '69' is strobogrammatic.")	
}

func Test_twopoint_ex2(t *testing.T) {
	assert := assert.New(t)
	
	ans := isStrobogrammatic("818")

	assert.Equal(true, ans, "The number '818' is strobogrammatic.")
}

func Test_twopoint_ex3(t *testing.T) {
	assert := assert.New(t)
	
	ans := isStrobogrammatic("962")

	assert.Equal(false, ans, "The number '96' is not strobogrammatic.")
}