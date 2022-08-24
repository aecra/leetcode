package problem309

import "testing"

// 309. Best Time to Buy and Sell Stock with Cooldown
// 309. 买卖股票的最佳时机含冷冻期
func TestMaxProfit(t *testing.T) {
	// test case 1
	got := maxProfit([]int{1, 2, 3, 0, 2})
	if got != 3 {
		t.Errorf("maxProfit() error %d.", got)
	}

	// test case 2
	got = maxProfit([]int{1})
	if got != 0 {
		t.Errorf("maxProfit() error %d.", got)
	}
}
