package problem402

func removeKdigits(num string, k int) string {
	// 去掉前导零
	for len(num) > 0 && num[0] == '0' {
		num = num[1:]
	}
	// 如果数字长度小于k，直接返回0
	if len(num) <= k {
		return "0"
	}
	preZero := false
	var res string
	for i := 0; i < k+1 && len(num) >= k+1; i++ {
		if num[i] == '0' {
			res = removeKdigits(num[i:], k-i)
			preZero = true
			break
		}
	}
	if !preZero {
		nums := 0
		var s StackByte
		for i := 0; i < len(num); i++ {
			for !s.empty() && s.top() > num[i] && nums < k {
				s.pop()
				nums++
			}
			s.push(num[i])
		}
		num = string(s.data)
		res = num[:len(num)-k+nums]
	}
	return res
}

type StackByte struct {
	data []byte
}

func (s *StackByte) push(val byte) {
	s.data = append(s.data, val)
}

func (s *StackByte) pop() byte {
	if len(s.data) == 0 {
		return 0
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

func (s *StackByte) top() byte {
	if len(s.data) == 0 {
		return 0
	}
	return s.data[len(s.data)-1]
}

func (s *StackByte) empty() bool {
	return len(s.data) == 0
}
