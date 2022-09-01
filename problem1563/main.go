package problem1563

func stoneGameV(stoneValue []int) int {
	if len(stoneValue) <= 1 {
		return 0
	}
	prefixSum := make([]int, len(stoneValue))
	prefixSum[0] = stoneValue[0]
	for i := 1; i < len(stoneValue); i++ {
		prefixSum[i] = prefixSum[i-1] + stoneValue[i]
	}
	dp := make([][]int, len(stoneValue)+1)
	for i := 0; i < len(stoneValue)+1; i++ {
		dp[i] = make([]int, len(stoneValue)+1)
		for j := 0; j < len(stoneValue)+1; j++ {
			dp[i][j] = -1
		}
	}
	return dfs(dp, stoneValue, prefixSum, 0, len(stoneValue))
}

func dfs(dp [][]int, stoneValue []int, prefixSum []int, start, end int) int {
	if end-start <= 1 {
		return 0
	}
	if dp[start][end] != -1 {
		return dp[start][end]
	}
	dp[start][end] = 0
	leftPrefix := 0
	rightPrefix := 0
	score := 0
	for i := start + 1; i < end; i++ {
		leftPrefix = prefixSum[i-1] - prefixSum[start] + stoneValue[start]
		rightPrefix = prefixSum[end-1] - prefixSum[i] + stoneValue[i]
		if leftPrefix > rightPrefix {
			score = rightPrefix + dfs(dp, stoneValue, prefixSum, i, end)
		} else if leftPrefix < rightPrefix {
			score = leftPrefix + dfs(dp, stoneValue, prefixSum, start, i)
		} else {
			score = max(leftPrefix+dfs(dp, stoneValue, prefixSum, start, i), rightPrefix+dfs(dp, stoneValue, prefixSum, i, end))
		}
		if score > dp[start][end] {
			dp[start][end] = score
		}
	}

	return dp[start][end]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
