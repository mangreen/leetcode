package q271_EncodeAndDecodeStrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_encodeAndDecode(t *testing.T) {
	assert := assert.New(t)

	input := "1/a2/ab3/abc"
	output := decode(input)
	ans := encode(output)

	assert.Equal("1/a2/ab3/abc", ans, `["a","ab","abc"] encode to be "1/a2/ab3/abc"`)
}