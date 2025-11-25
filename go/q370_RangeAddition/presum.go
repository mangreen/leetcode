package q370_RangeAddition

func getModifiedArray(length int, updates [][]int) []int {
    preSum := make([]int, length)
    
    // 應用差分數組技巧
    for _, update := range updates {
        start, end, inc := update[0], update[1], update[2]
        preSum[start] += inc
        if end+1 < length {
            preSum[end+1] -= inc
        }
    }
    
    // 計算前綴和以得到最終結果
    for i := 1; i < length; i++ {
        preSum[i] += preSum[i-1]
    }
    
    return preSum
}