package problem2235

import "testing"

func TestSum(t *testing.T) {
	// test case 1
	num1 := 12
	num2 := 5
	expected := 17
	actual := sum(num1, num2)
	if actual != expected {
		t.Errorf("TestSum failed, expected: %d, actual: %d", expected, actual)
	}

	// test case 2
	num1 = -10
	num2 = 4
	expected = -6
	actual = sum(num1, num2)
	if actual != expected {
		t.Errorf("TestSum failed, expected: %d, actual: %d", expected, actual)
	}
}
