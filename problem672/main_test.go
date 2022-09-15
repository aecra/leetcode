package problem672

import "testing"

func TestFlipLights(t *testing.T) {
	n := 1
	presses := 1
	expected := 2
	actual := flipLights(n, presses)
	if expected != actual {
		t.Errorf("n is %d, presses is %d, expected is %d, actual is %d", n, presses, expected, actual)
	}
	n = 2
	presses = 1
	expected = 3
	actual = flipLights(n, presses)
	if expected != actual {
		t.Errorf("n is %d, presses is %d, expected is %d, actual is %d", n, presses, expected, actual)
	}
	n = 3
	presses = 1
	expected = 4
	actual = flipLights(n, presses)
	if expected != actual {
		t.Errorf("n is %d, presses is %d, expected is %d, actual is %d", n, presses, expected, actual)
	}
}
