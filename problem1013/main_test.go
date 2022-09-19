package problem1013

import "testing"

func TestCanThreePartsEqualSum(t *testing.T) {
	// test case 1
	arr := []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}
	expected := true
	actual := canThreePartsEqualSum(arr)
	if actual != expected {
		t.Errorf("TestCanThreePartsEqualSum failed, expected: %v, actual: %v", expected, actual)
	}

	//  test case 2
	arr = []int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}
	expected = false
	actual = canThreePartsEqualSum(arr)
	if actual != expected {
		t.Errorf("TestCanThreePartsEqualSum failed, expected: %v, actual: %v", expected, actual)
	}

	// test case 3
	arr = []int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}
	expected = true
	actual = canThreePartsEqualSum(arr)
	if actual != expected {
		t.Errorf("TestCanThreePartsEqualSum failed, expected: %v, actual: %v", expected, actual)
	}

	// test case 4
	arr = []int{1, -1, 1, -1}
	expected = false
	actual = canThreePartsEqualSum(arr)
	if actual != expected {
		t.Errorf("TestCanThreePartsEqualSum failed, expected: %v, actual: %v", expected, actual)
	}

	// test case 5
	arr = []int{18, 12, -18, 18, -19, -1, 10, 10}
	expected = true
	actual = canThreePartsEqualSum(arr)
	if actual != expected {
		t.Errorf("TestCanThreePartsEqualSum failed, expected: %v, actual: %v", expected, actual)
	}

	// test case 6
	arr = []int{0, 0, 0, 0}
	expected = true
	actual = canThreePartsEqualSum(arr)
	if actual != expected {
		t.Errorf("TestCanThreePartsEqualSum failed, expected: %v, actual: %v", expected, actual)
	}
}
