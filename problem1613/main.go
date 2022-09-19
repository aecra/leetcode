package problem1613

import "sort"

func frequencySort(nums []int) []int {
	count := make(map[int]int)
	for _, i := range nums {
		count[i]++
	}
	sort.SliceStable(nums, func(i, j int) bool {
		return count[nums[i]] < count[nums[j]] || count[nums[i]] == count[nums[j]] && nums[i] > nums[j]
	})
	return nums
}
