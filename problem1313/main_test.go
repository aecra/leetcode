package problem1313

import (
	"reflect"
	"testing"
)

func TestDecompressRLElist(t *testing.T) {
	// test case 1
	input := []int{1, 2, 3, 4}
	expected := []int{2, 4, 4, 4}
	actual := decompressRLElist(input)
	if reflect.DeepEqual(expected, actual) == false {
		t.Errorf("expected %v, actual is %v", expected, actual)
	}

	// test case 2
	input = []int{1, 1, 2, 3}
	expected = []int{1, 3, 3}
	actual = decompressRLElist(input)
	if reflect.DeepEqual(expected, actual) == false {
		t.Errorf("expected %v, actual is %v", expected, actual)
	}
}
