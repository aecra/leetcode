package problem1354

import "testing"

func TestIsPossible(t *testing.T) {
	inputs := [][]int{
		{9, 3, 5},
		{1, 1, 1, 2},
		{8, 5},
		{9, 9, 9},
		{1, 1000000000},
	}
	expected := []bool{
		true,
		false,
		true,
		false,
		true,
	}
	for i, input := range inputs {
		if isPossible(input) != expected[i] {
			t.Errorf("Input %v Expected %v but got %v", input, expected[i], isPossible(input))
		}
	}
}
