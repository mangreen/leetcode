package q642_DesignSearchAutocompleteSystem

import (
	"sort"
)

type AutocompleteSystemTrie struct {
	root *trie
	cur  *trie
	prefix string
}

func ConstructorTrie(sentences []string, times []int) AutocompleteSystemTrie {
	root := &trie{
		children: make(map[byte]*trie),
		counts:   make(map[string]int),
	}

	for i, s := range sentences {
		root.insert(s, times[i])
	}

	return AutocompleteSystemTrie{
		root: root,
		cur:  root,
		prefix: "",
	}
}

func (a *AutocompleteSystemTrie) Input(c byte) []string {
	if c == '#' {
		// 結束輸入：將當前完整句子存入系統（頻率+1），重置狀態
		a.root.insert(a.prefix, 1)
		a.prefix = ""
		a.cur = a.root
		return []string{}
	}

	a.prefix += string(c)

	// 移動到當前前綴對應的節點
	if a.cur != nil {
		a.cur = a.cur.children[c]
	}

	// 若無對應節點，表示無匹配句子
	if a.cur == nil {
		return []string{}
	}

	// 收集當前節點的所有句子及其頻率，排序後返回前 3 個
	pairs := []pairTrie{}
	for s, cnt := range a.cur.counts {
		pairs = append(pairs, pairTrie{s, cnt})
	}

	// 排序邏輯：先比次數 (降序)，再比字典序 (升序)
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count == pairs[j].count {
			return pairs[i].sentence < pairs[j].sentence
		}

		return pairs[i].count > pairs[j].count
	})

	ans := []string{}
	for i := 0; i < len(pairs) && i < 3; i++ {
		ans = append(ans, pairs[i].sentence)
	}

	return ans
}

type trie struct {
	children map[byte]*trie // [26]*trie 會在遇到空格時導致 index out of range
	counts   map[string]int // 儲存經過該節點的句子及其頻率
}

func (t *trie) insert(s string, count int) {
	cur := t
	for _, r := range s {
		char := byte(r)

		if cur.children[char] == nil {
			cur.children[char] = &trie{
				children: make(map[byte]*trie),
				counts:   make(map[string]int),
			}
		}
		
		cur = cur.children[char]

		// 在每個節點更新該句子的頻率
		cur.counts[s] += count 
	}
}

type pairTrie struct {
	sentence string
	count    int
}