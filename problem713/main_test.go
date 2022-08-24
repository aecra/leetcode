package problem713

import "testing"

// 713. Subarray Product Less Than K
// 713. 乘积小于 K 的子数组
func TestNumSubarrayProductLessThanK(t *testing.T) {
	// test case 1
	got := numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100)
	if got != 8 {
		t.Errorf("numSubarrayProductLessThanK() error %d.", got)
	}
}
