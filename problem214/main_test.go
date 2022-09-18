package problem214

import "testing"

func TestShortestPalindrome(t *testing.T) {
	// test case 1
	s := "aacecaaa"
	expected := "aaacecaaa"
	actual := shortestPalindrome(s)
	if actual != expected {
		t.Errorf("test case 1 failed, expected: %s, actual: %s", expected, actual)
	}

	// test case 2
	s = "abcd"
	expected = "dcbabcd"
	actual = shortestPalindrome(s)
	if actual != expected {
		t.Errorf("test case 2 failed, expected: %s, actual: %s", expected, actual)
	}
}
