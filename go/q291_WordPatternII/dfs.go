package q291_WordPatternII

import "strings"

func wordPatternMatch(pattern string, str string) bool {
	m := make(map[byte]string)
	used := make(map[string]bool)

	return dfs(pattern, 0, str, 0, m, used)
}

func dfs(pat string, i int, str string, j int, m map[byte]string, used map[string]bool) bool {
	if i == len(pat) && j == len(str) {
		return true
	}

	if i >= len(pat) || j >= len(str) {
		return false
	}

	b := pat[i]
	if w, ok := m[b]; ok {
		if !strings.HasPrefix(str[j:], w) {
			return false
		}
	}

	for k:=j; k<len(str); k++ {
		sub := str[j:k+1]

		if used[sub] {
			continue
		}

		m[b] = sub
		used[sub] = true

		if dfs(pat, i+1, str, k+1, m, used) {
			return true
		}

		delete(m, b)
		delete(used, sub)
	}

	return false
}
