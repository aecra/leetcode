package problem84

import "testing"

// 84. Largest Rectangle in Histogram
// 84. 柱状图中最大的矩形
func TestLargestRectangleArea(t *testing.T) {
	// test case 1
	got := largestRectangleArea([]int{2, 1, 5, 6, 2, 3})
	if got != 10 {
		t.Errorf("largestRectangleArea() error. %d", got)
	}

	// test case 2
	got = largestRectangleArea([]int{4, 2, 0, 3, 2, 5})
	if got != 6 {
		t.Errorf("largestRectangleArea() error. %d", got)
	}

	// test case 3
	got = largestRectangleArea([]int{4})
	if got != 4 {
		t.Errorf("largestRectangleArea() error. %d", got)
	}
}
