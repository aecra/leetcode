package problem1894

func chalkReplacer(chalk []int, k int) int {
	sum := 0
	for _, c := range chalk {
		sum += c
	}
	k %= sum
	for i, c := range chalk {
		k -= c
		if k < 0 {
			return i
		}
	}
	return -1
}
