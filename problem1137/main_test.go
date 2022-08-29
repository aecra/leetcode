package problem1137

import "testing"

func TestTribonacci(t *testing.T) {
	// test case 1
	n := 4
	expected := 4
	actual := tribonacci(n)
	if actual != expected {
		t.Errorf("TestTribonacci failed, expected: %v, actual: %v", expected, actual)
	}

	// test case 2
	n = 25
	expected = 1389537
	actual = tribonacci(n)
	if actual != expected {
		t.Errorf("TestTribonacci failed, expected: %v, actual: %v", expected, actual)
	}

	// test case 3
	n = 34
	expected = 334745777
	actual = tribonacci(n)
	if actual != expected {
		t.Errorf("TestTribonacci failed, expected: %v, actual: %v", expected, actual)
	}
}
