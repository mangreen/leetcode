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

/*
★ 解題思路:
1. 用 Hash Map 來儲存 playerId -> score 的映射，方便快速查詢與更新。
2. 針對 top(K)，則有幾種做法：
  - 排序法：每次呼叫 top 時將所有分數取出並排序。時間複雜度 O(NlogN)。
  - 最小堆 (Min-Heap)：維護一個大小為 $K$ 的堆疊。時間複雜度 O(NlogK)。

ex.
1	addScore(1, 73)	⭢ m:{1: 73}
2	addScore(2, 56)	⭢ m:{1: 73, 2: 56}
3	addScore(3, 39)	⭢ m:{1: 73, 2: 56, 3: 39}
4	addScore(2, 20)	⭢ m:{1: 73, 2: 76, 3: 39}
5	top(2) ⭢ sort [76, 73, 39] ⭢ 76 + 73 = 149
6	reset(1)		⭢ m:{2: 76, 3: 39}
*/