package problem1700

import "testing"

func TestCountStudents(t *testing.T) {
	// test case 1
	students := []int{1, 1, 0, 0}
	sandwiches := []int{0, 1, 0, 1}
	if countStudents(students, sandwiches) != 0 {
		t.Error("Test case 1 failed")
	}

	// test case 2
	students = []int{1, 1, 1, 0, 0, 1}
	sandwiches = []int{1, 0, 0, 0, 1, 1}
	if countStudents(students, sandwiches) != 3 {
		t.Error("Test case 2 failed")
	}
}
