package problem658

import (
	"reflect"
	"testing"
)

// 658. Find K Closest Elements
// 658. 找到 K 个最接近的元素
func TestFindClosestElements(t *testing.T) {
	// test case 1
	got := findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3)
	if reflect.DeepEqual(got, []int{1, 2, 3, 4}) == false {
		t.Errorf("findClosestElements() error.")
	}

	// test case 2
	got = findClosestElements([]int{1, 2, 3, 4, 5}, 4, -1)
	if reflect.DeepEqual(got, []int{1, 2, 3, 4}) == false {
		t.Errorf("findClosestElements() error.")
	}
}
