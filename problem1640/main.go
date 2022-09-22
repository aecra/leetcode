package problem1640

func canFormArray(arr []int, pieces [][]int) bool {
	firstElement := make(map[int]int)
	for i, piece := range pieces {
		firstElement[piece[0]] = i
	}
	for i := 0; i < len(arr); {
		if _, ok := firstElement[arr[i]]; !ok {
			return false
		}
		piece := pieces[firstElement[arr[i]]]
		for j := 0; j < len(piece); j++ {
			if piece[j] != arr[i] || i >= len(arr) {
				return false
			}
			i++
		}
	}
	return true
}
