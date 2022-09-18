package problem5

import "testing"

func TestLongestPalindrome(t *testing.T) {
	// test case 1
	s := "babad"
	expected := "bab"
	actual := longestPalindrome(s)
	if actual != expected {
		t.Errorf("TestLongestPalindrome failed, expected %s, got %s", expected, actual)
	}

	// test case 2
	s = "cbbd"
	expected = "bb"
	actual = longestPalindrome(s)
	if actual != expected {
		t.Errorf("TestLongestPalindrome failed, expected %s, got %s", expected, actual)
	}
}
