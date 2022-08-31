package problem2391

import "testing"

func TestGarbageCollection(t *testing.T) {
	// test case 1
	garbage := []string{"G", "P", "GP", "GG"}
	travel := []int{2, 4, 3}
	expected := 21
	actual := garbageCollection(garbage, travel)
	if actual != expected {
		t.Errorf("TestGarbageCollection failed, expected: %d, actual: %d", expected, actual)
	}

	// test case 2
	garbage = []string{"MMM", "PGM", "GP"}
	travel = []int{3, 10}
	expected = 37
	actual = garbageCollection(garbage, travel)
	if actual != expected {
		t.Errorf("TestGarbageCollection failed, expected: %d, actual: %d", expected, actual)
	}
}
