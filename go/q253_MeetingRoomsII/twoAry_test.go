package q252_MeetingRooms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minMeetingRooms_1(t *testing.T) {
	assert := assert.New(t)

	input := [][]int{{0,30}, {5,10}, {15,20}}
	ans := minMeetingRooms(input)

	assert.Equal(2, ans, "[[0,30],[5,10],[15,20]] has overlap [0,30] & [5,10], so needs 2 rooms.")
}

func Test_minMeetingRooms_2(t *testing.T) {
	assert := assert.New(t)

	input := [][]int{{7,10}, {2,4}}
	ans := minMeetingRooms(input)

	assert.Equal(1, ans, "[[7,10],[2,4]] no overlap, so only need 1 room.")
}

func Test_minMeetingRooms_3(t *testing.T) {
	assert := assert.New(t)

	input := [][]int{{1,4}, {2,6}, {4,10}}
	ans := minMeetingRooms2(input)

	assert.Equal(2, ans, "[1,4] & [4, 10] can 1 room, and another 1 for [2,6].")
}

