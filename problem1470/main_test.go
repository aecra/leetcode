package problem1470

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	// test case 1
	nums := []int{2, 5, 1, 3, 4, 7}
	n := 3
	expected := []int{2, 3, 5, 4, 1, 7}
	actual := shuffle(nums, n)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("TestShuffle failed, expected: %v, actual: %v", expected, actual)
	}

	// test case 2
	nums = []int{1, 2, 3, 4, 4, 3, 2, 1}
	n = 4
	expected = []int{1, 4, 2, 3, 3, 2, 4, 1}
	actual = shuffle(nums, n)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("TestShuffle failed, expected: %v, actual: %v", expected, actual)
	}

	// test case 3
	nums = []int{1, 1, 2, 2}
	n = 2
	expected = []int{1, 2, 1, 2}
	actual = shuffle(nums, n)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("TestShuffle failed, expected: %v, actual: %v", expected, actual)
	}
}
