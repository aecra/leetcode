package problem1781

import (
	"testing"
)

func TestBeautySum(t *testing.T) {
	inputs := []string{"aabcb","aabcbaa"}
	expecteds := []int{5, 17}
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		expected := expecteds[i]
		actual := beautySum(input)
		if actual != expected {
			t.Errorf("beautySum(%v) != %v, actual %v", input, expected, actual)
		}
	}
}
