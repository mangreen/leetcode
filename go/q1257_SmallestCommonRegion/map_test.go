package q1257_SmallestCommonRegion

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_findSmallestRegion_1(t *testing.T) {
	assert := assert.New(t)

	regions := [][]string{{"Earth","North America","South America"},
		{"North America","United States","Canada"},
		{"United States","New York","Boston"},
		{"Canada","Ontario","Quebec"},
		{"South America","Brazil"}}
	
	r1 := "Quebec"
	r2 := "New York"

	ans := findSmallestRegion(regions, r1, r2)

	assert.Equal(ans, "North America", "North America is the LCA of Quebec & New York")
}

func Test_findSmallestRegion_2(t *testing.T) {
	assert := assert.New(t)

	regions := [][]string{{"Earth","North America","South America"},
		{"North America","United States","Canada"},
		{"United States","New York","Boston"},
		{"Canada","Ontario","Quebec"},
		{"South America","Brazil"}}
	
	r1 := "Canada"
	r2 := "South America"

	ans := findSmallestRegion(regions, r1, r2)

	assert.Equal(ans, "Earth", "Earth is the LCA of Canada & South America")
}
