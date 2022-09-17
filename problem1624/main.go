package problem1624

func maxLengthBetweenEqualCharacters(s string) int {
	max := -1
	firstIndex := make(map[rune]int)
	for i, c := range s {
		if _, ok := firstIndex[c]; !ok {
			firstIndex[c] = i
		} else {
			if i-firstIndex[c]-1 > max {
				max = i - firstIndex[c] - 1
			}
		}
	}
	return max
}
