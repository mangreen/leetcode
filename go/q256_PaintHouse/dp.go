package q256_PaintHouse

import ()

func minCost(costs [][]int) int {
    if len(costs)==0 || len(costs[0]) ==0 {
        return 0
    }

    dp := make([][]int, len(costs))
    copy(dp, costs)

    for i:=1; i<len(dp); i++ {
        dp[i][0] += min(dp[i-1][1], dp[i-1][2])
        dp[i][1] += min(dp[i-1][0], dp[i-1][2])
        dp[i][2] += min(dp[i-1][1], dp[i-1][0])
    }

    return min(min(dp[len(dp)-1][0], dp[len(dp)-1][1]), dp[len(dp)-1][2])
}
/*
ex1.
costs=[[17,2,17],[16,16,5],[14,3,19]]
dp=[[17,2,17],[16,16,5],[14,3,19]]

          min(16+17, 16+17)
i=2                |
dp=[[17,2,17],[18,33,7],[14,3,19]]
               /      \
min(16+2, 16+17)      min(5+17, 5+2)

                    min(3+18, 3+7)
i=2                          |
dp=[[17,2,17],[18,33,7],[21,10,37]]
                         /      \
           min(14+33, 14+7)      min(19+18, 19+33)

ans=min(21, 10, 37)=10
*/