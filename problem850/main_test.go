package problem850

import "testing"

func TestRectangleArea(t *testing.T) {
	// test case 1
	rectangles := [][]int{{0, 0, 2, 2}, {1, 0, 2, 3}, {1, 0, 3, 1}}
	expected := 6
	actual := rectangleArea(rectangles)
	if actual != expected {
		t.Errorf("expected %d, but got %d", expected, actual)
	}

	// test case 2
	rectangles = [][]int{{0, 0, 1000000000, 1000000000}}
	expected = 49
	actual = rectangleArea(rectangles)
	if actual != expected {
		t.Errorf("expected %d, but got %d", expected, actual)
	}
}
