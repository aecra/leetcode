package problem1827

import "testing"

// 1827. Minimum Operations to Make the Array Increasing
// 1827. 最少操作使数组递增
func TestMinOperations(t *testing.T) {
	// test case 1
	got := minOperations([]int{1, 1, 1})
	if got != 3 {
		t.Errorf("minOperations() error %d.", got)
	}

	// test case 2
	got = minOperations([]int{1, 5, 2, 4, 1})
	if got != 14 {
		t.Errorf("minOperations() error %d.", got)
	}
}
