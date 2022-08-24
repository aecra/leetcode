package problem1827

func minOperations(nums []int) int {
	operation := 0
	base := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > base {
			base = nums[i]
		} else {
			operation += base - nums[i] + 1
			base += 1
		}
	}
	return operation
}
