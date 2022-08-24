package problem402

import "testing"

// 402. Remove K Digits
// 402. 移掉 K 位数字
func TestRemoveKdigits(t *testing.T) {
	// test case 1
	got := removeKdigits("1432219", 3)
	if got != "1219" {
		t.Errorf("removeKdigits() error %s.", got)
	}

	// test case 2
	got = removeKdigits("10200", 1)
	if got != "200" {
		t.Errorf("removeKdigits() error %s.", got)
	}

	// test case 3
	got = removeKdigits("10", 1)
	if got != "0" {
		t.Errorf("removeKdigits() error %s.", got)
	}

	// test case 4
	got = removeKdigits("1140004278", 4)
	if got != "278" {
		t.Errorf("removeKdigits() error %s.", got)
	}

	// test case 5
	got = removeKdigits("278", 2)
	if got != "2" {
		t.Errorf("removeKdigits() error %s.", got)
	}
}
