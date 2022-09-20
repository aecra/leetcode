package problem1658

import "sort"

func minOperations(nums []int, x int) int {
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}
	if prefixSum[len(nums)-1] < x {
		return -1
	}
	if prefixSum[len(nums)-1] == x {
		return len(nums)
	}

	minOperation := len(nums) + 1
	left := prefixSum[len(nums)-1] - x
	for i := 0; i < len(nums); i++ {
		j := sort.Search(len(prefixSum), func(j int) bool {
			return prefixSum[j] >= prefixSum[i]-nums[i]+left
		})
		if j > len(prefixSum)-1 {
			continue
		}
		if prefixSum[j] == prefixSum[i]-nums[i]+left {
			minOperation = min(minOperation, len(nums)-(j-i+1))
		}
	}

	if minOperation == len(nums)+1 {
		return -1
	}
	return minOperation
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
