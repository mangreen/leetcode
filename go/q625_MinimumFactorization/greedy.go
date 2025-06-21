package q625_MinimumFactorization

import (
	"math"
)


func smallestFactorization(a int) int {
	if a < 10 {
		return a
	}

	ans := 0
	step := 1

	for i:=9; i>=2; i-- {
		for a%i == 0 {
			ans = step*i + ans
			step *= 10
			a /= i

			if a == 1 {
				if ans > math.MaxInt32 {
					return 0
				}

				return ans
			}
		}
	}

	return 0
}
/*
i=9
648%9=0
ans=9
step=10
a=648/9=72

72%9=0
ans=10*9+9=99
step=100
a=72/9=8

8/9 != 0

i=8
8%8=0
ans=100*8+99=899 ⭠ ans
step=1000
a=8/8=1
*/