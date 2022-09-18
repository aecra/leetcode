package problem5

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var start, end int
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		maxlen := max(len1, len2)
		if maxlen > end-start+1 {
			start = i - (maxlen-1)/2
			end = i + maxlen/2
		}
	}
	return s[start : end+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func expandAroundCenter(s string, left, right int) int {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}
