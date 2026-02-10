package q248_StrobogrammaticNumberIII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dfs_ex1(t *testing.T) {
	assert := assert.New(t)

	ans := strobogrammaticInRange("50", "100")

	assert.Equal(3, ans, "The strobogrammatic numbers in the range [50, 100] are '69', '88', and '96'.")
}

func Test_dfs_ex2(t *testing.T) {
	assert := assert.New(t)

	ans := strobogrammaticInRange("0", "0")

	assert.Equal(1, ans, "The only strobogrammatic number in the range [0, 0] is '0'.")
}

func Test_dfs_ex3(t *testing.T) {
	assert := assert.New(t)

	ans := strobogrammaticInRange("10", "15")

	assert.Equal(1, ans, "The only strobogrammatic number in the range [10, 15] is '11'.")
}
