package problem1781

func beautySum(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		for j := i + 2; j <= len(s); j++ {
			sum += beauty(s[i:j])
		}
	}
	return sum
}

func beauty(s string) int {
	count := [26]int{}
	for _, c := range s {
		count[c-'a']++
	}
	min, max := 1<<31, 0
	for _, c := range count {
		if c == 0 {
			continue
		}
		if c < min {
			min = c
		}
		if c > max {
			max = c
		}
	}
	return max - min
}
