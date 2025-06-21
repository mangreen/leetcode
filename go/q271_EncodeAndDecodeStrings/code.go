package q271_EncodeAndDecodeStrings

import (
	"strconv"
	"strings"
)

func encode(strs []string) string {
	ans := ""

	for _, s := range strs {
		ans += strconv.Itoa(len(s)) + "/" + s
	}

	return ans
}

func decode(s string) []string {
	ans := []string{}

	for s != "" {
		i := strings.Index(s, "/")
		L, _ := strconv.Atoi(string(s[i-1]))
		ans = append(ans, s[i+1:i+1+L])
		s = s[i+1+L:]
	}

	return ans
}
