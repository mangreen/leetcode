package q1151_MinimumSwapsToGroupAll1sTogether

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mostVisitedPattern_1(t *testing.T) {
	assert := assert.New(t)

	username := []string{"joe","joe","joe","james","james","james","james","mary","mary","mary"}
	timestamp := []int{1,2,3,4,5,6,7,8,9,10}
	website := []string{"home","about","career","home","cart","maps","home","home","about","career"}
	ans := mostVisitedPattern(username, timestamp, website)

	assert.Equal([]string{"home", "about", "career"}, ans, "The pattern 'home', 'about', 'career' is visited by Joe and Mary, making it the most visited pattern.")
}

func Test_mostVisitedPattern_2(t *testing.T) {
	assert := assert.New(t)

	username := []string{"ua","ua","ua","ub","ub","ub"}
	timestamp := []int{1,2,3,4,5,6}
	website := []string{"a","b","a","a","b","c"}
	ans := mostVisitedPattern(username, timestamp, website)

	assert.Equal([]string{"a", "b", "a"}, ans, "The pattern 'a', 'b', 'a' is visited by user 'ub', making it the most visited pattern.")
}
