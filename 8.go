package leetcode

func myAtoi(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			s = s[i:]
			break
		}
	}
	if len(s) == 0 {
		return 0
	}
	sign := 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		sign = 1
		s = s[1:]
	} else if s[0] > '9' || s[0] < '0' {
		return 0
	}
	for i := 0; i < len(s); i++ {
		if s[i] > '9' || s[i] < '0' {
			s = s[:i]
			break
		}
	}
	var b int64 = 0
	for v := range s {
		b = b*10 + int64(s[v]-'0')
		if sign == 1 {
			if b >= 2147483647 {
				return 2147483647
			}
		} else {
			if b >= 2147483648 {
				return -2147483648
			}
		}
	}

	return sign * int(b)
}
