package offerII64

func twoSum(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i < j; {
		if numbers[i]+numbers[j] == target {
			return []int{i, j}
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}
	return []int{-1, -1}
}
