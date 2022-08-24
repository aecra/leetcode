package problem84

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	var s1 stack
	var s2 stack
	maxAera := 0
	heights = append(heights, 0)
	for i := 0; i < len(heights); i++ {
		if !s1.empty() && heights[s1.top()] == heights[i] {
			continue
		}
		if s1.empty() || heights[s1.top()] < heights[i] {
			s1.push(i)
		} else {
			for !s1.empty() {
				maxAera = max(maxAera, heights[s1.top()]*(i-s1.top()))
				s2.push(s1.pop())
			}
			for !s2.empty() {
				if heights[i] >= heights[s2.top()] {
					s1.push(s2.pop())
				} else {
					heights[s2.top()] = heights[i]
					s1.push(s2.pop())
					break
				}
			}
			s2.clear()
		}
	}
	return maxAera
}

type stack struct {
	data []int
}

func (s *stack) push(val int) {
	s.data = append(s.data, val)
}

func (s *stack) pop() int {
	if len(s.data) == 0 {
		return 0
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

func (s *stack) top() int {
	if len(s.data) == 0 {
		return 0
	}
	return s.data[len(s.data)-1]
}

func (s *stack) empty() bool {
	return len(s.data) == 0
}

func (s *stack) clear() {
	s.data = []int{}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
