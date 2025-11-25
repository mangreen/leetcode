package q269_AlienDictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_alienOrder_1(t *testing.T) {
	assert := assert.New(t)

	input := []string{"wrt","wrf","er","ett","rftt"}
	ans := alienOrder(input)

	assert.Equal(ans, "wertf","")
}

func Test_alienOrder_2(t *testing.T) {
	assert := assert.New(t)

	input := []string{"z","x"}
	ans := alienOrder(input)

	assert.Equal(ans, "zx","")
}

func Test_alienOrder_3(t *testing.T) {
	assert := assert.New(t)

	input := []string{"z","x","z"}
	ans := alienOrder(input)

	assert.Equal(ans, "","")
}

func Test_alienOrder_4(t *testing.T) {
	assert := assert.New(t)

	input := []string{"abc","ab"}
	ans := alienOrder(input)

	assert.Equal(ans, "","")
}

func Test_alienOrder_5(t *testing.T) {
	assert := assert.New(t)

	input := []string{"a", "b", "ca", "cc"}
	ans := alienOrder(input)

	assert.Equal(ans, "abc","abc 或 acb")
}