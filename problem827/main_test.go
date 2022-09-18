package problem827

import "testing"

func TestLargestIsland(t *testing.T) {
	// test case 1
	grid := [][]int{
		{1, 0},
		{0, 1},
	}
	expected := 3
	actual := largestIsland(grid)
	if actual != expected {
		t.Errorf("TestLargestIsland failed, expected %d, got %d", expected, actual)
	}

	// test case 2
	grid = [][]int{
		{1, 1},
		{1, 0},
	}
	expected = 4
	actual = largestIsland(grid)
	if actual != expected {
		t.Errorf("TestLargestIsland failed, expected %d, got %d", expected, actual)
	}

	// test case 3
	grid = [][]int{
		{1, 1},
		{1, 1},
	}
	expected = 4
	actual = largestIsland(grid)
	if actual != expected {
		t.Errorf("TestLargestIsland failed, expected %d, got %d", expected, actual)
	}
}
