package q660_Remove9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newInteger_9_str(t *testing.T) {
	assert := assert.New(t)

	ans := newInteger(9)

	assert.Equal(10, ans, "1,2,3,4,5,6,7,8,x,10 ⭢ 10 is the 9th number without 9")
}

func Test_newInteger_19_str(t *testing.T) {
	assert := assert.New(t)

	ans := newInteger(19)

	assert.Equal(21, ans, "1,2,3,4,5,6,7,8,x,10,11,12,13,14,15,16,17,18,x,20,21 ⭢ 21 is the 19th number without 9")
}
