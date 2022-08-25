package problem81

import "testing"

func TestSearch(t *testing.T) {
	// test case 1
	input := []int{2, 5, 6, 0, 0, 1, 2}
	target := 0
	expected := true
	actual := search(input, target)
	if actual != expected {
		t.Errorf("TestSearch: expected %v, actual %v", expected, actual)
	}
	// test case 2
	input = []int{2, 5, 6, 0, 0, 1, 2}
	target = 3
	expected = false
	actual = search(input, target)
	if actual != expected {
		t.Errorf("TestSearch: expected %v, actual %v", expected, actual)
	}
	// test case 3
	input = []int{1}
	target = 0
	expected = false
	actual = search(input, target)
	if actual != expected {
		t.Errorf("TestSearch: expected %v, actual %v", expected, actual)
	}
	// test case 4
	input = []int{3, 1}
	target = 0
	expected = false
	actual = search(input, target)
	if actual != expected {
		t.Errorf("TestSearch: expected %v, actual %v", expected, actual)
	}
	// test case 4
	input = []int{3, 1}
	target = 3
	expected = true
	actual = search(input, target)
	if actual != expected {
		t.Errorf("TestSearch: expected %v, actual %v", expected, actual)
	}
	// test case 5
	input = []int{8, 9, 2, 3, 4}
	target = 9
	expected = true
	actual = search(input, target)
	if actual != expected {
		t.Errorf("TestSearch: expected %v, actual %v", expected, actual)
	}
	// test case 6
	input = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}
	target = 2
	expected = true
	actual = search(input, target)
	if actual != expected {
		t.Errorf("TestSearch: expected %v, actual %v", expected, actual)
	}
}
