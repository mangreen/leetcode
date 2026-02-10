package q186_ReverseWordsInAStringII

func reverseWords(s []byte) {
	N := len(s)

	// 先反轉整個字串
	reverse(s, 0, N-1)

	// 再反轉每個單字
	start := 0
	for i := 0; i <= N; i++ {
		if i == N || s[i] == ' ' {
			reverse(s, start, i-1)
			start = i + 1
		}	
	}
}

func reverse(s []byte, l, r int) {
	for l < r {
		s[l], s[r] = s[r], s[l]	
		l++
		r--
	}
}