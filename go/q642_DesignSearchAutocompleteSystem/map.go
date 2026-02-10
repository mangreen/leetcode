package q642_DesignSearchAutocompleteSystem

import (
	"sort"
	"strings"
)

// 暴力解法
type AutocompleteSystem struct {
	m      map[string]int // 儲存所有句子及其頻率
	prefix string         // 當前輸入的前綴
	cur    map[string]int // 當前候選句子
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	m := make(map[string]int)
	for i, s := range sentences {
		m[s] = times[i]
	}

	return AutocompleteSystem{
		m,
		"",
		m,
	}
}

// Input 模擬使用者輸入字元，返回匹配的前 3 個句子
func (a *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		a.m[string(a.prefix)]++
		a.prefix = ""
		a.cur = a.m
		return []string{}
	}

	a.prefix += string(c)

	ps := []pair{}

	// 篩選出匹配當前前綴的句子
	for s, t := range a.cur {
		// 判斷句子是否以當前前綴開頭
		if strings.HasPrefix(s, a.prefix) {
			ps = append(ps, pair{
				s,
				t,
			})
		} else {
			delete(a.cur, s)
		}
	}

	// 排序邏輯：先比次數 (降序)，再比字典序 (升序)
	sort.Slice(ps, func(i, j int) bool {
		a := ps[i]
		b := ps[j]
		return a.t > b.t || (a.t == b.t && a.s < b.s)
	})

	ans := []string{}

	// 取前 3 個結果
	for i := range 3 {
		if len(ps) < i+1 {
			break
		}

		ans = append(ans, ps[i].s)
	}

	return ans
}

type pair struct {
	s string
	t int
}
