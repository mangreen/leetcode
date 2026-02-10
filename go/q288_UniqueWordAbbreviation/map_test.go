package q288_UniqueWordAbbreviation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_map_ex1(t *testing.T) {
	assert := assert.New(t)

	dict := []string{"deer", "door", "cake", "card"}
	v := Constructor(dict)

	/* 
		m:{"d2r": "deer"}
		m:{"d2r": ""}
		m:{"d2r": "", "c2e": "cake"}
		m:{"d2r": "", "c2e": "cake", "c2d": "card"}
	*/
	assert.Equal(false, v.IsUnique("deer"))
	assert.Equal(true, v.IsUnique("cart"))
	assert.Equal(false, v.IsUnique("cane"))
	assert.Equal(true, v.IsUnique("make"))
	assert.Equal(true, v.IsUnique("cake"))
}

func Test_map_ex2(t *testing.T) {
	assert := assert.New(t)

	dict := []string{"bay", "buy", "diy", "day"}
	v := Constructor(dict)

	/* 
		m:{"b1y": "bay"}
		m:{"b1y": ""}
		m:{"b1y": "", "d1y": "diy"}
		m:{"b1y": "", "d1y": ""}
	*/
	assert.Equal(false, v.IsUnique("bay"))
	assert.Equal(false, v.IsUnique("buy"))
	assert.Equal(false, v.IsUnique("diy"))
	assert.Equal(false, v.IsUnique("day"))
	assert.Equal(true, v.IsUnique("cat"))
}
