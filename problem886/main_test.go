package problem886

import "testing"

func TestPossibleBipartition(t *testing.T) {
	// test case 1
	n := 4
	dislikes := [][]int{{1, 2}, {1, 3}, {2, 4}}
	if possibleBipartition(n, dislikes) != true {
		t.Errorf("expected true, got false")
	}

	// test case 2
	n = 3
	dislikes = [][]int{{1, 2}, {1, 3}, {2, 3}}
	if possibleBipartition(n, dislikes) != false {
		t.Errorf("expected false, got true")
	}

	// test case 3
	n = 5
	dislikes = [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}
	if possibleBipartition(n, dislikes) != false {
		t.Errorf("expected false, got true")
	}
}
