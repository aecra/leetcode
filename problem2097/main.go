package problem2097

import "fmt"

func validArrangement(pairs [][]int) [][]int {
	indegree := make(map[int]int)
	outdegree := make(map[int]int)
	firstElementIndex := make(map[int][]int)
	for i, pair := range pairs {
		firstElementIndex[pair[0]] = append(firstElementIndex[pair[0]], i)
		indegree[pair[0]]++
		outdegree[pair[1]]++
	}
	// find the start node
	start := pairs[0][0]
	for k, v := range indegree {
		if outdegree[k] < v {
			start = k
			break
		}
	}
	fmt.Println("first step", start)
	res := make([][]int, 0, len(pairs))
	used := make([]bool, len(pairs))
	var dfs func(int) bool
	dfs = func(node int) bool {
		if len(res) == len(pairs) {
			return true
		}
		for _, i := range firstElementIndex[node] {
			if used[i] {
				continue
			}
			used[i] = true
			res = append(res, pairs[i])
			if dfs(pairs[i][1]) {
				return true
			}
			res = res[:len(res)-1]
			used[i] = false
		}
		return false
	}
	dfs(start)
	return res
}
