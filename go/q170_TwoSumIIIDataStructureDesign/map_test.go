package q170_TwoSumIIIDataStructureDesign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_map_ex1(t *testing.T) {
	assert := assert.New(t)

	ts := Constructor()
	ts.Add(1)
	ts.Add(3)
	ts.Add(5)

	assert.Equal(true, ts.Find(4))
}

func Test_map_ex2(t *testing.T) {
	assert := assert.New(t)

	ts := Constructor()
	ts.Add(1)
	ts.Add(3)
	ts.Add(5)

	assert.Equal(false, ts.Find(7))
}