package problem1563

import "testing"

func TestStoneGameV(t *testing.T) {
	// test case 1
	input := []int{6, 2, 3, 4, 5, 5}
	expected := 18
	actual := stoneGameV(input)
	if actual != expected {
		t.Errorf("Test Case 1: expected = %d, actual = %d", expected, actual)
	}

	// test case 2
	input = []int{7, 7, 7, 7, 7, 7, 7}
	expected = 28
	actual = stoneGameV(input)
	if actual != expected {
		t.Errorf("Test Case 2: expected = %d, actual = %d", expected, actual)
	}

	// test case 3
	input = []int{4}
	expected = 0
	actual = stoneGameV(input)
	if actual != expected {
		t.Errorf("Test Case 3: expected = %d, actual = %d", expected, actual)
	}
}
