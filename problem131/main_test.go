package problem131

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	// test case 1
	input := "aab"
	expected := [][]string{
		{"a", "a", "b"},
		{"aa", "b"},
	}
	actual := partition(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("TestPartition: expected %v, actual %v", expected, actual)
	}
	// test case 2
	input = "a"
	expected = [][]string{
		{"a"},
	}
	actual = partition(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("TestPartition: expected %v, actual %v", expected, actual)
	}
	// test case 3
	input = "bb"
	expected = [][]string{
		{"b", "b"},
		{"bb"},
	}
	actual = partition(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("TestPartition: expected %v, actual %v", expected, actual)
	}
}
