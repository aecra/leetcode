package leetcode

func findClosestElements(arr []int, k int, x int) []int {
	point := 0
	diff := 1<<31 - 1
	for i := 0; i < len(arr); i++ {
		if abs(arr[i]-x) < diff {
			point = i
			diff = abs(arr[i] - x)
		}
	}
	left, right := point, point
	k--
	for (left > 0 || right < len(arr)-1) && k > 0 {
		k--
		if left == 0 {
			right++
			continue
		}
		if right == len(arr)-1 {
			left--
			continue
		}
		if abs(arr[left-1]-x) <= abs(arr[right+1]-x) {
			left--
		} else {
			right++
		}

	}
	return arr[left : right+1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
