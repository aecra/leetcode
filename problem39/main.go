package problem39

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	result = dfs(result, candidates, []int{}, target)
	return result
}

func dfs(result [][]int, candidates []int, arr []int, target int) [][]int {
	for i := 0; i < len(candidates); i++ {
		if candidates[i] > target {
			continue
		} else if candidates[i] == target {
			arr = append(arr, candidates[i])
			// copy arr to append to result
			new_arr := make([]int, len(arr))
			copy(new_arr, arr)
			result = append(result, new_arr)
			arr = arr[:len(arr)-1]
		} else {
			arr = append(arr, candidates[i])
			result = dfs(result, candidates[i:], arr, target-candidates[i])
			arr = arr[:len(arr)-1]
		}
	}
	return result
}
