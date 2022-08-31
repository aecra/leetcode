package problem789

func escapeGhosts(ghosts [][]int, target []int) bool {
	step := abs(target[0]) + abs(target[1])
	for _, ghost := range ghosts {
		if abs(ghost[0]-target[0])+abs(ghost[1]-target[1]) <= step {
			return false
		}
	}
	return true
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
