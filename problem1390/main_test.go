package problem1390

import "testing"

func TestSumFourDivisors(t *testing.T) {
	// test case 1
	nums := []int{21, 4, 7}
	expected := 32
	actual := sumFourDivisors(nums)
	if actual != expected {
		t.Errorf("TestSumFourDivisors failed, expected: %d, actual: %d", expected, actual)
	}

	// test case 2
	nums = []int{21, 21}
	expected = 64
	actual = sumFourDivisors(nums)
	if actual != expected {
		t.Errorf("TestSumFourDivisors failed, expected: %d, actual: %d", expected, actual)
	}

	// test case 2
	nums = []int{1, 2, 3, 4, 5}
	expected = 0
	actual = sumFourDivisors(nums)
	if actual != expected {
		t.Errorf("TestSumFourDivisors failed, expected: %d, actual: %d", expected, actual)
	}
}
