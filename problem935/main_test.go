package problem935

import "testing"

func TestKnightDialer(t *testing.T) {
	// test case 1
	n := 1
	expected := 10
	actual := knightDialer(n)
	if actual != expected {
		t.Errorf("TestKnightDialer failed, expected:%v, actual:%v", expected, actual)
	}

	// test case 2
	n = 2
	expected = 20
	actual = knightDialer(n)
	if actual != expected {
		t.Errorf("TestKnightDialer failed, expected:%v, actual:%v", expected, actual)
	}

	// test case 3
	n = 3131
	expected = 136006598
	actual = knightDialer(n)
	if actual != expected {
		t.Errorf("TestKnightDialer failed, expected:%v, actual:%v", expected, actual)
	}
}
