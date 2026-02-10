package q248_StrobogrammaticNumberIII

func strobogrammaticInRange(low string, high string) int {
	cnt := 0

	// 枚举长度
	for L := len(low); L <= len(high); L++ {
		DFS(low, high, "", L, &cnt)
		DFS(low, high, "0", L, &cnt)
		DFS(low, high, "1", L, &cnt)
		DFS(low, high, "8", L, &cnt)
	}	

	return cnt
}


func DFS(low, high, cur string, maxL int, cnt *int) {
	if len(cur) > maxL {
		return
	}

	if len(cur) == maxL {
		// 不能以0开头
		if len(cur) > 1 && cur[0] == '0' {
			return
		}

		// 超出范围
		if (len(cur) == len(low) && cur < low) || (len(cur) == len(high) && cur > high) {
			return
		}	

		*cnt++
		return
	}

	// 递归构造
	DFS(low, high, "0"+cur+"0", maxL, cnt)
	DFS(low, high, "1"+cur+"1", maxL, cnt)
	DFS(low, high, "6"+cur+"9", maxL, cnt)
	DFS(low, high, "8"+cur+"8", maxL, cnt)
	DFS(low, high, "9"+cur+"6", maxL, cnt)
}
/*
ex1.
low="50", high="100"
             ↱ ["000","101","609","808","906"]
L=2 ⭢ ["", "0", "1", "8"] ⭢ ["080", "181","689","888","986"]
        ⭣         ↳ ["010", "111","619","818","916"]
["11","69","88","96"] 
count=3 只有 "69","88","96" 在范围内


ex2.
low="0", high="0"
L=1 ⭢ ["0","1","8"]
count=1

ex3.
low="10", high="15"
L=2 ⭢ ["11","69","88","96"]
count=1
*/