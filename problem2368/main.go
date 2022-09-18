package problem2368

func reachableNodes(n int, edges [][]int, restricted []int) int {
	adjacencyList := make([][]int, n)
	for _, edge := range edges {
		adjacencyList[edge[0]] = append(adjacencyList[edge[0]], edge[1])
		adjacencyList[edge[1]] = append(adjacencyList[edge[1]], edge[0])
	}
	visited := make([]bool, n)
	for _, r := range restricted {
		visited[r] = true
	}
	reachableNodesNum := 1
	visited[0] = true
	stack := []int{0}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, neighbor := range adjacencyList[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				reachableNodesNum++
				stack = append(stack, neighbor)
			}
		}
	}
	return reachableNodesNum
}
