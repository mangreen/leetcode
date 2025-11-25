package q1087_BraceExpansion

import (
	"sort"
)

func expand(s string) []string {
	options := [][]byte{}

	// 字串分組 ex. "{a,b}c{d,e}f" ⭢ [[a,b],[c],[d,e],f]
	for i := 0; i < len(s); i++ {
		if s[i] == '{' {
			i++

			opt := []byte{}

			for s[i] != '}' {
				if s[i] != ',' {
					opt = append(opt, s[i])
				}

				i++
			}

			sort.Slice(opt, func(i, j int) bool {
				return opt[i] < opt[j]
			})

			options = append(options, opt)
		} else {
			options = append(options, []byte{s[i]})
		}
	}

	ans := []string{}

	DFS(options, 0, []byte{}, &ans)

	return ans
}

func DFS(options [][]byte, idx int, cur []byte, ans *[]string) {
	if idx == len(options) {
		*ans = append(*ans, string(cur))
		return
	}

	for _, b := range options[idx] {
		DFS(options, idx+1, append(cur, b), ans)
	}
}
