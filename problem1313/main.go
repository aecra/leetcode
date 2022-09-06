package problem1313

func decompressRLElist(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums)/2; i++ {
		for j := 0; j < nums[2*i]; j++ {
			res = append(res, nums[2*i+1])
		}
	}
	return res
}
