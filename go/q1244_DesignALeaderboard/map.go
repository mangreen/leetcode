package q1244_DesignALeaderboard

import (
	"sort"
)

type Leaderboard struct {
	m map[int]int // playerId -> score
}

func Constructor() Leaderboard {
	return Leaderboard{
		m: make(map[int]int),
	}
}

func (this *Leaderboard) AddScore(playerId int, score int) {
	this.m[playerId] += score
}

func (this *Leaderboard) Top(K int) int {
	// Extract scores into a slice for sorting
	playerScores := make([]int, 0, len(this.m))
	for _, s := range this.m {
		playerScores = append(playerScores, s)
	}

	// Sort in descending order
	sort.Slice(playerScores, func(i, j int) bool {
		return playerScores[i] > playerScores[j]
	})

	// Sum the top K scores
	sum := 0
	for _, s := range playerScores[:K] {
		sum += s
	}
    
	return sum
}

func (this *Leaderboard) Reset(playerId int) {
	delete(this.m, playerId) // Or set score to 0: this.scores[playerId] = 0
}

/**
 * Your Leaderboard object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddScore(playerId,score);
 * param_2 := obj.Top(K);
 * obj.Reset(playerId);
 */