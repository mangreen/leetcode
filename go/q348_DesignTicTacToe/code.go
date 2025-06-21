package q348_DesignTicTacToe

type TicTacToe struct {
	n   int
	cnt [][]int
}

func Constructor(n int) TicTacToe {
	cnt := make([][]int, 2)
	for i := range cnt {
		cnt[i] = make([]int, (n<<1)+2)
	}

	return TicTacToe{n, cnt}
}

func (ttt *TicTacToe) Move(row int, col int, player int) int {
	cur := ttt.cnt[player-1]
	cur[row]++
	cur[ttt.n + col]++

	if row == col {
		cur[ttt.n<<1]++
	}
	
	if row+col == ttt.n-1 {
		cur[ttt.n<<1|1]++
	}
	
	if cur[row] == ttt.n || cur[ttt.n+col] == ttt.n || cur[ttt.n<<1] == ttt.n || cur[ttt.n<<1|1] == ttt.n {
		return player
	}

	return 0
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
