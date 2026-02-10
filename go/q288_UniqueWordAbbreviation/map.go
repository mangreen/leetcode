package q288_UniqueWordAbbreviation

import "strconv"

type ValidWordAbbr struct {
	dic map[string]string
}

func Constructor(dictionary []string) ValidWordAbbr {
	m := make(map[string]string)
	for _, s := range dictionary {
		abbr := getAbbr(s)

		// 如果縮寫已存在且對應的原始單字不同，設為空字串表示不唯一
		if val, ok := m[abbr]; ok {
			if val != s {
				m[abbr] = ""
			}
		} else {
			m[abbr] = s
		}
	}

	return ValidWordAbbr{m}
}

func (this *ValidWordAbbr) IsUnique(word string) bool {
	abbr := getAbbr(word)
	val, ok := this.dic[abbr]

	// 1. 縮寫沒出現過 -> Unique
    // 2. 縮寫出現過，且對應的原始單字就是自己 -> Unique
	return !ok || val == word
}

func getAbbr(s string) string {
	N := len(s)
	if N <= 2 {
		return s
	}

	return string(s[0]) + strconv.Itoa(N-2) + string(s[N-1])
}
