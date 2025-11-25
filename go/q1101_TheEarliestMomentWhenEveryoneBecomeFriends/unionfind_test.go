package q1101_TheEarliestMomentWhenEveryoneBecomeFriends

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_unionfind_1(t *testing.T) {
	assert := assert.New(t)

	logs := [][]int{{20190101, 0, 1},{20190104, 3, 4},{20190107, 2, 3},{20190211, 1, 5},{20190224, 2, 4},{20190301, 0, 3},{20190312, 1, 2},{20190322, 4, 5}} 
	n := 6

	ans := earliestAcq(logs, n)

	assert.Equal(20190301, ans, "20190301")
}

func Test_unionfind_2(t *testing.T) {
	assert := assert.New(t)

	logs := [][]int{{0,2,0},{1,0,1},{3,0,3},{4,1,2},{7,3,1}}
	n := 4

	ans := earliestAcq(logs, n)

	assert.Equal(3, ans, "3")
}
