package problem1894

import "testing"

func TestChalkReplacer(t *testing.T) {
	// test case 1
	chalk := []int{5, 1, 5}
	k := 22
	expected := 0
	actual := chalkReplacer(chalk, k)
	if actual != expected {
		t.Errorf("chalkReplacer(%v, %v) = %v; expected %v", chalk, k, actual, expected)
	}

	// test case 2
	chalk = []int{3, 4, 1, 2}
	k = 25
	expected = 1
	actual = chalkReplacer(chalk, k)
	if actual != expected {
		t.Errorf("chalkReplacer(%v, %v) = %v; expected %v", chalk, k, actual, expected)
	}
}
