package problem647

import "testing"

func TestCountSubstrings(t *testing.T) {
	// test case 1
	s := "abc"
	expected := 3
	actual := countSubstrings(s)
	if actual != expected {
		t.Errorf("TestCountSubstrings failed, expected: %v, actual: %v", expected, actual)
	}

	// test case 2
	s = "aaa"
	expected = 6
	actual = countSubstrings(s)
	if actual != expected {
		t.Errorf("TestCountSubstrings failed, expected: %v, actual: %v", expected, actual)
	}
}
