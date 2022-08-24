package problem557

func reverseWords(s string) string {
	res := []byte{}
	for i := 0; i < len(s); {
		start := i
		for i < len(s) && s[i] != ' ' {
			i++
		}

		for point := start; point < i; point++ {
			res = append(res, s[start+i-1-point])
		}

		for i < len(s) && s[i] == ' ' {
			i++
			res = append(res, ' ')
		}
	}
	return string(res)
}
