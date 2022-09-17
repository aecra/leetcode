package lcp55

import "testing"

func TestGetMinimumTime(t *testing.T) {
	// test case 1
	time := []int{2, 3, 2}
	fruits := [][]int{
		{0, 2},
		{1, 4},
		{2, 1},
	}
	limit := 3
	expected := 10
	actual := getMinimumTime(time, fruits, limit)
	if actual != expected {
		t.Errorf("expected %v, but got %v", expected, actual)
	}

	// test case 2
	time = []int{1}
	fruits = [][]int{
		{0, 3},
		{0, 5},
	}
	limit = 2
	expected = 5
	actual = getMinimumTime(time, fruits, limit)
	if actual != expected {
		t.Errorf("expected %v, but got %v", expected, actual)
	}
}
