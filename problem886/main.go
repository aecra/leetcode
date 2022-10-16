package problem886

func possibleBipartition(n int, dislikes [][]int) bool {
	color := make([]int, n+1)
	graph := make([][]int, n+1)
	for _, d := range dislikes {
		graph[d[0]] = append(graph[d[0]], d[1])
		graph[d[1]] = append(graph[d[1]], d[0])
	}
	var dfs func(int, int) bool
	dfs = func(n, c int) bool {
		color[n] = c
		for _, next := range graph[n] {
			if color[next] == c {
				return false
			}
			if color[next] == 0 && !dfs(next, -c) {
				return false
			}
		}
		return true
	}
	for i := 1; i <= n; i++ {
		if color[i] == 0 && !dfs(i, 1) {
			return false
		}
	}
	return true
}