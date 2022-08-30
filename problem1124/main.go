package problem1124

func longestWPI(hours []int) int {
	if len(hours) == 0 {
		return 0
	}
	tired := make([]int, len(hours))
	prefixSum := make([]int, len(hours))
	for i := 0; i < len(hours); i++ {
		if hours[i] > 8 {
			tired[i] = 1
		}
		if i == 0 {
			prefixSum[i] = tired[i]
		} else {
			prefixSum[i] = prefixSum[i-1] + tired[i]
		}
	}
	maxLength := 0
	for i := 0; i < len(hours); i++ {
		for j := i + maxLength; j < len(hours); j++ {
			if (prefixSum[j]-prefixSum[i]+tired[i])*2 > j-i+1 {
				maxLength = j - i + 1
			}
		}
	}
	return maxLength
}
