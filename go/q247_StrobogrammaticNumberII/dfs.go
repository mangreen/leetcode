package q247_StrobogrammaticNumberII

func findStrobogrammatic(n int) []string {
	return dfs(n, n)
}

func dfs(n, maxN int) []string {
	if n == 0 {
		return []string{""}
	}	

	if n == 1 {
		return []string{"0", "1", "8"}
	}
	
	ans := []string{}
	for _, v := range dfs(n-2, maxN) {
		// 0 只有在不是最外层时才可以使用
		if n != maxN {
			ans = append(ans, "0"+v+"0")
		}

		ans = append(ans, "1"+v+"1")
		ans = append(ans, "6"+v+"9")
		ans = append(ans, "8"+v+"8")
		ans = append(ans, "9"+v+"6")
	}

	return ans
}

/*
ex1.
n=3 
⭣
n-2=1 ⭢ ["0","1","8"] 
⭣
["101","609","808","906","111","619","818","916","181","689","888","986"]

ex2.
n=2
⭣
n-2=0 ⭢ [""]
⭣
["11","69","88","96"]
*/
