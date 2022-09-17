package problem1624

import "testing"

func TestMaxLengthBetweenEqualCharacters(t *testing.T) {
	// test case 1
	s := "aa"
	expected := 0
	actual := maxLengthBetweenEqualCharacters(s)
	if actual != expected {
		t.Errorf("expected %d, but got %d", expected, actual)
	}

	// test case 2
	s = "abca"
	expected = 2
	actual = maxLengthBetweenEqualCharacters(s)
	if actual != expected {
		t.Errorf("expected %d, but got %d", expected, actual)
	}

	// test case 3
	s = "cbzxy"
	expected = -1
	actual = maxLengthBetweenEqualCharacters(s)
	if actual != expected {
		t.Errorf("expected %d, but got %d", expected, actual)
	}

	// test case 4
	s = "cabbac"
	expected = 4
	actual = maxLengthBetweenEqualCharacters(s)
	if actual != expected {
		t.Errorf("expected %d, but got %d", expected, actual)
	}
}
