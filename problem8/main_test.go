package problem8

import "testing"

// 8. String to Integer (atoi)
// 8. 字符串转换整数 (atoi)
func TestMyAtoi(t *testing.T) {
	// test case 1
	got := myAtoi("42")
	if got != 42 {
		t.Errorf("myAtoi() error: 42 -> %d", got)
	}

	// test case 2
	got = myAtoi("   -42")
	if got != -42 {
		t.Errorf("myAtoi() error: -42 -> %d", got)
	}

	// test case 3
	got = myAtoi("4193 with words")
	if got != 4193 {
		t.Errorf("myAtoi() error: 4193 -> %d", got)
	}
}
