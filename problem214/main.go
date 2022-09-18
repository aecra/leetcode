package problem214

func shortestPalindrome(s string) string {
	// this is similler as longestPalindrome
	// find the longest palindrome substring which start from index 0
	// then reverse the rest of the string and append it to the original string
	// time usage: 284ms
	// memory usage: 7.8MB
	if len(s) < 2 {
		return s
	}
	var index int
	var double bool
	for i := 0; i < len(s); i++ {
		res1 := expandAroundCenter(s, i, i)
		res2 := expandAroundCenter(s, i, i+1)
		if res1 {
			index, double = i, false
		}
		if res2 {
			index, double = i, true
		}
	}
	var res string
	if double {
		res = stringReverse(s[(index+1)*2:])
	} else {
		res = stringReverse(s[(index+1)*2-1:])
	}
	return res + s
}

func stringReverse(s string) string {
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}

func expandAroundCenter(s string, left, right int) bool {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return L == -1
}
