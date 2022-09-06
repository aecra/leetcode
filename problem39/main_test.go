package problem39

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	// test case 1
	candidates := []int{2, 3, 6, 7}
	target := 7
	expected := [][]int{
		{2, 2, 3},
		{7},
	}
	actual := combinationSum(candidates, target)
	if reflect.DeepEqual(expected, actual) == false {
		t.Errorf("expected %v, got %v", expected, actual)
	}

	// test case 2
	candidates = []int{2, 3, 5}
	target = 8
	expected = [][]int{
		{2, 2, 2, 2},
		{2, 3, 3},
		{3, 5},
	}
	actual = combinationSum(candidates, target)
	if reflect.DeepEqual(expected, actual) == false {
		t.Errorf("expected %v, got %v", expected, actual)
	}

	// test case 3
	candidates = []int{2}
	target = 1
	expected = [][]int{}
	actual = combinationSum(candidates, target)
	if reflect.DeepEqual(expected, actual) == false {
		t.Errorf("expected %v, got %v", expected, actual)
	}

}
