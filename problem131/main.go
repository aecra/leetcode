package problem131

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	if len(s) == 1 {
		return [][]string{{s}}
	}
	var ret [][]string
	for i := 1; i < len(s); i++ {
		if isPalindrome(s[:i]) {
			for _, r := range partition(s[i:]) {
				ret = append(ret, append([]string{s[:i]}, r...))
			}
		}
	}
	if isPalindrome(s) {
		ret = append(ret, []string{s})
	}
	return ret
}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	for i := 0; i < len(s)-i-1; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
