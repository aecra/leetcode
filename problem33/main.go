package problem33

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	} else if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	// first we need find the index
	startIndex := getStartIndex(nums)
	// find the target
	left, right := startIndex, startIndex+len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[(mid+len(nums))%len(nums)] == target {
			return mid % len(nums)
		} else if nums[(mid+len(nums))%len(nums)] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func getStartIndex(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right && !isStartIndex(nums, left) {
		mid := (left + right) / 2
		if nums[mid] > nums[left] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			left++
		}
	}
	return left
}

func isStartIndex(nums []int, index int) bool {
	if index < 0 || index >= len(nums) {
		return false
	}
	if index == 0 {
		return nums[index] < nums[index+1] && nums[index] < nums[len(nums)-1]
	}
	if index == len(nums)-1 {
		return nums[index] < nums[index-1] && nums[index] < nums[0]
	}
	return nums[index] < nums[index+1] && nums[index] < nums[index-1]
}
