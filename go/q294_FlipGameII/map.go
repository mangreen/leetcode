package q294_FlipGameII

func canWin(currentState string) bool {
    // 使用 map 進行記憶化搜尋 (Memoization) 優化效能
    memo := make(map[string]bool)
    return helper(currentState, memo)
}

func helper(curS string, memo map[string]bool) bool {
	// 檢查是否已經計算過當前狀態
	if val, ok := memo[curS]; ok {
		return val
	}

	for i := range curS[:len(curS)-1] {
		if curS[i:i+2] == "++" {
			nextState := curS[:i] + "--" + curS[i+2:]

			// 如果對手在 nextState 無法贏，則我方勝利
			if !helper(nextState, memo) {
				memo[curS] = true
				
				return true
			}
		}
	}

	memo[curS] = false
	
	return false
}
/*
實例解說： 輸入 currentState = "++++"
1. 玩家 A 翻轉成 +--+。
2. 此時 玩家 B 面對 +--+，發現沒有任何連續的 ++ 可以翻轉。
3. 玩家 B 輸了，因此 玩家 A 在 ++++ 狀態下是「必勝」的。
*/