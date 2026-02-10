package q246_StrobogrammaticNumber

func isStrobogrammatic(num string) bool {
	dict := map[byte]byte{
		'0': '0',
		'1': '1',
		'6': '9',
		'8': '8',
		'9': '6',
	}

	l, r := 0, len(num)-1
	for l <= r {
		if v, ok := dict[num[l]]; !ok || v != num[r] {
			return false
		}

		l++
		r--
	}

	return true
}