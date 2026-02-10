package q660_Remove9

func newInteger2(n int) int {
	ans, base := 0, 1
    for n > 0 {
        ans += (n % 9) * base
        n /= 9
        base *= 10
    }
    return ans
}
/*
ex1: n = 19
ans=0, base=1
19÷9=2...1 → ans=1*1=1, base=10
2÷9=0...2 → ans=1+2*10=21, base=100
ans=21

├──────── 9 ────────┼──────────── 9 ─────────────┼1┤
1,2,3,4,5,6,7,8,x,10,11,12,13,14,15,16,17,18,x,20,21
*/

