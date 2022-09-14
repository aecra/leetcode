package offerII64

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	// test case 1
	numbers := []int{1, 2, 4, 6, 10}
	target := 8
	expected := []int{1, 3}
	actual := twoSum(numbers, target)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test case 1 failed, actual: %v, expected: %v", actual, expected)
	}

	// test case 2
	numbers = []int{2, 3, 4}
	target = 6
	expected = []int{0, 2}
	actual = twoSum(numbers, target)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test case 2 failed, actual: %v, expected: %v", actual, expected)
	}

	// test case 3
	numbers = []int{-1, 0}
	target = -1
	expected = []int{0, 1}
	actual = twoSum(numbers, target)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test case 3 failed, actual: %v, expected: %v", actual, expected)
	}
}
