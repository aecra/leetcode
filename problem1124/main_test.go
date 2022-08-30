package problem1124

import "testing"

func TestLongestWPI(t *testing.T) {
	// test case 1
	hours := []int{9, 9, 6, 0, 6, 6, 9}
	expected := 3
	actual := longestWPI(hours)
	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
