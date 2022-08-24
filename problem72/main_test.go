package problem72

import "testing"

// 72. Edit Distance
// 72. 编辑距离
func TestMinDistance(t *testing.T) {
	// test case 1
	got := minDistance("horse", "ros")
	if got != 3 {
		t.Errorf("minDistance() error %d.", got)
	}

	// test case 2
	got = minDistance("intention", "execution")
	if got != 5 {
		t.Errorf("minDistance() error %d.", got)
	}
}
