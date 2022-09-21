package problem854

import "testing"

func TestKSimilarity(t *testing.T) {
	tests := []struct {
		s1 string
		s2 string
		k  int
	}{
		{"ab", "ba", 1},
		{"abc", "bca", 2},
		{"abac", "baca", 2},
		{"aabc", "abca", 2},
	}
	for _, test := range tests {
		if got := kSimilarity(test.s1, test.s2); got != test.k {
			t.Errorf("KSimilarity(%q, %q) = %v; want %v", test.s1, test.s2, got, test.k)
		}
	}
}
