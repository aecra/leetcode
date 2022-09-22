package problem1640

import "testing"

func TestCanFormArrary(t *testing.T) {
	tests := []struct {
		arr    []int
		pieces [][]int
		want   bool
	}{
		{[]int{15, 88}, [][]int{{88}, {15}}, true},
		{[]int{49, 18, 16}, [][]int{{16, 18, 49}}, false},
		{[]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}}, true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := canFormArray(tt.arr, tt.pieces); got != tt.want {
				t.Errorf("canFromArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
