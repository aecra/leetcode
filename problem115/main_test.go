package problem115

import "testing"

func TestNumDistinct(t *testing.T) {
	// test case 1
	input := "rabbbit"
	pattern := "rabbit"
	expected := 3
	actual := numDistinct(input, pattern)
	if actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}

	// test case 2
	input = "babgbag"
	pattern = "bag"
	expected = 5
	actual = numDistinct(input, pattern)
	if actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}
