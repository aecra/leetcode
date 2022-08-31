package problem789

import "testing"

func TestEscapeGhosts(t *testing.T) {
	// test case 1
	ghosts := [][]int{{1, 0}, {0, 3}}
	target := []int{0, 1}
	expected := true
	actual := escapeGhosts(ghosts, target)
	if actual != expected {
		t.Errorf("TestEscapeGhosts failed, expected: %t, actual: %t", expected, actual)
	}

	// test case 2
	ghosts = [][]int{{1, 0}}
	target = []int{2, 0}
	expected = false
	actual = escapeGhosts(ghosts, target)
	if actual != expected {
		t.Errorf("TestEscapeGhosts failed, expected: %t, actual: %t", expected, actual)
	}

	// test case 3
	ghosts = [][]int{{2, 0}}
	target = []int{1, 0}
	expected = false
	actual = escapeGhosts(ghosts, target)
	if actual != expected {
		t.Errorf("TestEscapeGhosts failed, expected: %t, actual: %t", expected, actual)
	}
}
