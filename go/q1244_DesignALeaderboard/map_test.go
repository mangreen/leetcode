package q1244_DesignALeaderboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_leaderboard(t *testing.T) {
	assert := assert.New(t)

	leaderboard := Constructor()

	leaderboard.AddScore(1,73) // leaderboard = [[1,73]]
	leaderboard.AddScore(2,56) // leaderboard = [[1,73],[2,56]]
	leaderboard.AddScore(3,39) // leaderboard = [[1,73],[2,56],[3,39]]
	leaderboard.AddScore(4,51) // leaderboard = [[1,73],[2,56],[3,39],[4,51]]
	leaderboard.AddScore(5,4) // leaderboard = [[1,73],[2,56],[3,39],[4,51],[5,4]]
	
	assert.Equal(leaderboard.Top(1), 73, "It should return 73")

	leaderboard.Reset(1)
	leaderboard.Reset(2)

	leaderboard.AddScore(2,51)

	assert.Equal(leaderboard.Top(3), 141, "It should return 141 = 51 + 51 + 39")
}
