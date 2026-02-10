package q360_SortTransformedArray

func sortTransformedArray(nums []int, a int, b int, c int) []int {
    N := len(nums)
    ans := make([]int, N)
    left, right := 0, N-1

    if a >= 0 {
        // 開口向上: 較大值在兩端，從後面開始放
        idx := N - 1

        for left <= right {
            leftVal := f(a, b, c, nums[left])
            rightVal := f(a, b, c, nums[right]) 

            if leftVal >= rightVal {
                ans[idx] = leftVal
                left++
            } else {
                ans[idx] = rightVal
                right--
            }

            idx--
        }
    } else {
        // 開口向下: 較小值在兩端，從前面開始放
        idx := 0

        for left <= right {
            leftVal := f(a, b, c, nums[left])
            rightVal := f(a, b, c, nums[right]) 

            if leftVal <= rightVal {
                ans[idx] = leftVal
                left++
            } else {
                ans[idx] = rightVal
                right--
            }

            idx++
        }
    }

    return ans
}

func f(a, b, c, x int) int {
	return a*x*x + b*x + c
}
/*
不能直接用 sort.Ints()？ 因為那樣複雜度是 O(N log N)
題目要求 O(N)，所以要利用 nums 已經排序以及拋物線圖形的特性。

f(x) = ax^2 + bx + c 微笑曲線

1. 當 a >= 0 時，拋物線開口向上
兩端的值較大，所以從兩端往中間找大的，放在結果陣列的最後面 (idx=N-1)。
|      |
\      /
 ↘ _ ↗

2. 當 a < 0 時，拋物線開口向下
中間的值較大，所以從兩端往中間找小的，放在結果陣列的最前面 (idx=0)。
    _
  ↙  ↘
 /     \
|       |

使用雙指標法，從陣列的兩端開始比較，根據 a 的正負決定將較大或較小的值放入結果陣列中，直到所有元素處理完畢。
*/