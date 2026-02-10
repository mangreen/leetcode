package q286_WallsAndGates

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_bfs_ex1(t *testing.T) {
	assert := assert.New(t)

	rooms := [][]int{
		{2147483647, -1, 0, 2147483647},
		{2147483647, 2147483647, 2147483647, -1},
		{2147483647, -1, 2147483647, -1},
		{0, -1, 2147483647, 2147483647},
	}

	wallsAndGatesBFS(rooms)

	expected := [][]int{
		{3, -1, 0, 1},
		{2, 2, 1, -1},
		{1, -1, 2, -1},
		{0, -1, 3, 4},
	}

	if assert.NotNil(rooms) {
		assert.Equal(expected, rooms, "The rooms after wallsAndGates DFS traversal should match the expected distances")
	}
}

func Test_bfs_ex2(t *testing.T) {
	assert := assert.New(t)

	rooms := [][]int{{-1}}

	wallsAndGatesBFS(rooms)

	expected := [][]int{{-1}}

	if assert.NotNil(rooms) {
		assert.Equal(expected, rooms, "The rooms after wallsAndGates DFS traversal should match the expected distances")
	}
}
