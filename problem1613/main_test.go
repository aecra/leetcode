package problem1613

import (
	"reflect"
	"testing"
)

func TestFrequencySort(t *testing.T) {
	// test case 1
	nums := []int{1, 1, 2, 2, 2, 3}
	expected := []int{3, 1, 1, 2, 2, 2}
	actual := frequencySort(nums)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("TestFrequencySort failed, expected: %v, actual: %v", expected, actual)
	}

	// test case 2
	nums = []int{2, 3, 1, 3, 2}
	expected = []int{1, 3, 3, 2, 2}
	actual = frequencySort(nums)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("TestFrequencySort failed, expected: %v, actual: %v", expected, actual)
	}

	// test case 3
	nums = []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}
	expected = []int{5, -1, 4, 4, -6, -6, 1, 1, 1}
	actual = frequencySort(nums)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("TestFrequencySort failed, expected: %v, actual: %v", expected, actual)
	}
}
