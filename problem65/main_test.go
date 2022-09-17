package problem65

import "testing"

func TestIsnumber(t *testing.T) {
	// test case 1
	s := "0"
	expected := true
	actual := isNumber(s)
	if actual != expected {
		t.Errorf("expected %v, but got %v", expected, actual)
	}

	// test case 2
	s = "e"
	expected = false
	actual = isNumber(s)
	if actual != expected {
		t.Errorf("expected %v, but got %v", expected, actual)
	}

	// test case 3
	s = "."
	expected = false
	actual = isNumber(s)
	if actual != expected {
		t.Errorf("expected %v, but got %v", expected, actual)
	}
}
