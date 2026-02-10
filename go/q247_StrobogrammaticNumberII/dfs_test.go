package q247_StrobogrammaticNumberII

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dfs_ex1(t *testing.T) {
	assert := assert.New(t)

	ans := findStrobogrammatic(2)

	assert.Equal([]string{"11", "69", "88", "96"}, ans, "The strobogrammatic numbers of length 2 are '11', '69', '88', and '96'.")
}

func Test_dfs_ex2(t *testing.T) {
	assert := assert.New(t)

	ans := findStrobogrammatic(3)
	sort.Strings(ans)

	expected := []string{"101", "609", "808", "906", "111", "619", "818", "916", "181", "689", "888", "986"}
	sort.Strings(expected)

	assert.Equal(expected, ans, "The strobogrammatic numbers of length 3 are '101', '111', '181', '609', '619', '689', '808', '818', '888', '906', '916', and '986'.")
}

func Test_dfs_ex3(t *testing.T) {
	assert := assert.New(t)

	ans := findStrobogrammatic(1)

	assert.Equal([]string{"0", "1", "8"}, ans, "The strobogrammatic numbers of length 1 are '0', '1', and '8'.")
}
