package problem347

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	// test case 1
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	expected := []int{1, 2}
	actual := topKFrequent(nums, k)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}

	// test case 2
	nums = []int{1}
	k = 1
	expected = []int{1}
	actual = topKFrequent(nums, k)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
