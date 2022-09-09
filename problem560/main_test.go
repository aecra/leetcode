package problem560

import "testing"

func TestSubarraySum(t *testing.T) {
	// test cases 1
	nums := []int{1, 1, 1}
	k := 2
	expected := 2
	actual := subarraySum(nums, k)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}

	// test casse 2
	nums = []int{1, 2, 3}
	k = 3
	expected = 2
	actual = subarraySum(nums, k)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
