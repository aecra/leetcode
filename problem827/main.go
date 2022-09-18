package problem827

func dfs(grid, areaColor [][]int, color, i, j int) int {
	areaColor[i][j] = color
	area := 1
	n := len(grid)
	if i > 0 && grid[i-1][j] == 1 && areaColor[i-1][j] == -1 {
		areaColor[i-1][j] = color
		area += dfs(grid, areaColor, color, i-1, j)
	}
	if i < n-1 && grid[i+1][j] == 1 && areaColor[i+1][j] == -1 {
		areaColor[i+1][j] = color
		area += dfs(grid, areaColor, color, i+1, j)
	}
	if j > 0 && grid[i][j-1] == 1 && areaColor[i][j-1] == -1 {
		areaColor[i][j-1] = color
		area += dfs(grid, areaColor, color, i, j-1)
	}
	if j < n-1 && grid[i][j+1] == 1 && areaColor[i][j+1] == -1 {
		areaColor[i][j+1] = color
		area += dfs(grid, areaColor, color, i, j+1)
	}
	return area
}

func largestIsland(grid [][]int) int {
	n := len(grid)
	areaColor := make([][]int, n)
	for i := 0; i < n; i++ {
		areaColor[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			areaColor[i][j] = -1
		}
	}
	islandArea := make(map[int]int)
	maxIsland := 0
	// color grid and calculate area of each island
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && areaColor[i][j] == -1 {
				area := dfs(grid, areaColor, n*i+j, i, j)
				islandArea[n*i+j] = area
				if area > maxIsland {
					maxIsland = area
				}
			}
		}
	}
	// find largest island after adding one water
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				neighbours := make(map[int]bool)
				if i > 0 && grid[i-1][j] == 1 {
					neighbours[areaColor[i-1][j]] = true
				}
				if i < n-1 && grid[i+1][j] == 1 {
					neighbours[areaColor[i+1][j]] = true
				}
				if j > 0 && grid[i][j-1] == 1 {
					neighbours[areaColor[i][j-1]] = true
				}
				if j < n-1 && grid[i][j+1] == 1 {
					neighbours[areaColor[i][j+1]] = true
				}
				area := 1
				for k := range neighbours {
					area += islandArea[k]
				}
				if area > maxIsland {
					maxIsland = area
				}
			}
		}
	}
	return maxIsland
}
