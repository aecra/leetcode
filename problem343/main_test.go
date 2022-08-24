package problem343

import "testing"

// 343. Integer Break
// 343. 整数拆分
func TestIntegerBreak(t *testing.T) {
	// test case 1
	got := integerBreak(2)
	if got != 1 {
		t.Errorf("integerBreak() error %d.", got)
	}

	// test case 2
	got = integerBreak(3)
	if got != 2 {
		t.Errorf("integerBreak() error %d.", got)
	}

	// test case 3
	got = integerBreak(9)
	if got != 27 {
		t.Errorf("integerBreak() error %d.", got)
	}

	// test case 4
	got = integerBreak(10)
	if got != 36 {
		t.Errorf("integerBreak() error %d.", got)
	}
}
