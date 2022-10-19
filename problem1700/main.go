package problem1700

func countStudents(students []int, sandwiches []int) int {
	nums := make([]int, 2)
	for _, s := range students {
		nums[s]++
	}
	for {
		if len(sandwiches) == 0 {
			return 0
		}
		if nums[sandwiches[0]] == 0 {
			return len(students)
		}
		if students[0] == sandwiches[0] {
			nums[students[0]]--
			students = students[1:]
			sandwiches = sandwiches[1:]
		} else {
			students = append(students[1:], students[0])
		}
	}
}