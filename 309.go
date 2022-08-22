package leetcode

func maxProfit309(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dp := make([][3]int, len(prices))
	dp[0][0] = -prices[0] // 买入
	dp[0][1] = 0          // 卖出
	dp[0][2] = 0          // 冷冻期
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max309(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = max309(dp[i-1][1], dp[i-1][0]+prices[i])
		dp[i][2] = dp[i-1][1]
	}
	maxProfit := max309(max309(dp[len(prices)-1][0], dp[len(prices)-1][1]), dp[len(prices)-1][2])
	return maxProfit
}

func max309(a, b int) int {
	if a > b {
		return a
	}
	return b
}
