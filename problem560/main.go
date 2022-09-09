package problem560

func subarraySum(nums []int, k int) int {
	prefix := 0
	count := 0
	m := map[int]int{0: 1}
	for _, num := range nums {
		prefix += num
		if _, ok := m[prefix-k]; ok {
			count += m[prefix-k]
		}
		m[prefix]++
	}
	return count
}
