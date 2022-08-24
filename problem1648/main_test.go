package problem1648

import "testing"

// 1648. Maximum Product of Two Elements in an Array
// 1648. 销售价值减少的颜色球
func TestMaxProfit(t *testing.T) {
	// test case 1
	got := maxProfit([]int{2, 5}, 4)
	if got != 14 {
		t.Errorf("maxProfit() error: 14 -> %d", got)
	}

	// test case 2
	got = maxProfit([]int{1000000000}, 1000000000)
	if got != 21 {
		t.Errorf("maxProfit() error: 21 -> %d", got)
	}

	// test case 3
	got = maxProfit([]int{2, 8, 4, 10, 6}, 30)
	if got != 125 {
		t.Errorf("maxProfit() error: 125 -> %d", got)
	}

	// test case 4
	got = maxProfit([]int{680754224, 895956841, 524658639, 3161416, 992716630, 7365440, 582714485, 422256708, 332815744, 269407767}, 707568720)
	if got != 6150955 {
		t.Errorf("maxProfit() error: 6150955 -> %d", got)
	}
}
