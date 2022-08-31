package problem539

import "testing"

func TestFindMinDifference(t *testing.T) {
	// test case 1
	timePoints := []string{"23:59", "00:00"}
	expected := 1
	actual := findMinDifference(timePoints)
	if actual != expected {
		t.Errorf("TestFindMinDifference failed, expected: %d, actual: %d", expected, actual)
	}

	// test case 2
	timePoints = []string{"00:00", "23:59", "00:00"}
	expected = 0
	actual = findMinDifference(timePoints)
	if actual != expected {
		t.Errorf("TestFindMinDifference failed, expected: %d, actual: %d", expected, actual)
	}
}
