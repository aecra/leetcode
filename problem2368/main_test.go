package problem2368

import "testing"

func TestReachableNodes(t *testing.T) {
	// test case 1
	n := 7
	edges := [][]int{
		{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6},
	}
	restricated := []int{4, 5}
	expected := 4
	actual := reachableNodes(n, edges, restricated)
	if actual != expected {
		t.Errorf("TestReachableNodes failed, expected %d, got %d", expected, actual)
	}

	// test case 2
	n = 7
	edges = [][]int{
		{0, 1}, {0, 2}, {0, 5}, {0, 4}, {3, 2}, {6, 5},
	}
	restricated = []int{4, 2, 1}
	expected = 3
	actual = reachableNodes(n, edges, restricated)
	if actual != expected {
		t.Errorf("TestReachableNodes failed, expected %d, got %d", expected, actual)
	}
}
