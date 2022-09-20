package problem1658

import "testing"

func TestMinOperations(t *testing.T) {
	// test case 1
	nums := []int{1, 1, 4, 2, 3}
	x := 5
	expected := 2
	actual := minOperations(nums, x)
	if actual != expected {
		t.Errorf("TestMinOperations failed, expected: %v, actual: %v", expected, actual)
	}

	// test case 2
	nums = []int{5, 6, 7, 8, 9}
	x = 4
	expected = -1
	actual = minOperations(nums, x)
	if actual != expected {
		t.Errorf("TestMinOperations failed, expected: %v, actual: %v", expected, actual)
	}

	// test case 3
	nums = []int{3, 2, 20, 1, 1, 3}
	x = 10
	expected = 5
	actual = minOperations(nums, x)
	if actual != expected {
		t.Errorf("TestMinOperations failed, expected: %v, actual: %v", expected, actual)
	}
}
