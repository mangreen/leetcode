package q660_Remove9

import "strconv"

func newInteger(n int) int {
	res := ""

    for n > 0 {
        // 將 n % 9 轉成字串，並「接在前面」
        res = strconv.Itoa(n%9) + res
        n /= 9
    }

    // 將最後組好的字串轉回整數
    ans, _ := strconv.Atoi(res)
    return ans
}
/*
ex1: n = 19
19÷9=2...1 → res="1"
2÷9=0...2 → res="2"+"1"="21"
ans=21

├──────── 9 ────────┼──────────── 9 ─────────────┼1┤
1,2,3,4,5,6,7,8,x,10,11,12,13,14,15,16,17,18,x,20,21
*/

