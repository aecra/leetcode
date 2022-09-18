package problem516

import "testing"

func TestLongestPalindromeSubseq(t *testing.T) {
	// test case 1
	s := "bbbab"
	expected := 4
	actual := longestPalindromeSubseq(s)
	if expected != actual {
		t.Errorf("test case 1 failed, expected: %d, actual: %d", expected, actual)
	}

	// test case 2
	s = "cbbd"
	expected = 2
	actual = longestPalindromeSubseq(s)
	if expected != actual {
		t.Errorf("test case 2 failed, expected: %d, actual: %d", expected, actual)
	}
}
