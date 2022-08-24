package problem2178

import (
	"reflect"
	"testing"
)

// 2178. Maximum Split of Positive Even Integers
// 2178. 拆分成最多数目的正偶数之和
func TestMaximumEvenSplit(t *testing.T) {
	// test case 1
	got := maximumEvenSplit(12)
	if reflect.DeepEqual(got, []int64{2, 4, 6}) == false {
		t.Errorf("maximumEvenSplit() error.")
	}

	// test case 2
	got = maximumEvenSplit(7)
	if len(got) != 0 {
		t.Errorf("maximumEvenSplit() error.")
	}

	// test case 3
	got = maximumEvenSplit(28)
	if reflect.DeepEqual(got, []int64{2, 4, 6, 16}) == false {
		t.Errorf("maximumEvenSplit() error.")
	}
}
