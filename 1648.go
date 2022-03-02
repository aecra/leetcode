package leetcode

/**
 * 将数量按颜色排序，从高到低依次减小 orders
 */
func maxProfit(inventory []int, orders int) int {
	quickSort(inventory)
	mod := 1000000007
	res := 0
	point := len(inventory) - 1
	var drop int
	height := inventory[point]
	if point == 0 {
		drop = height
	} else {
		drop = height - inventory[point-1]
	}

	for orders > 0 {
		for drop == 0 {
			if point == 0 {
				drop = height
				break
			} else {
				drop = height - inventory[point-1]
				if drop == 0 {
					point--
				}
			}
		}

		// 优化
		width := len(inventory) - point
		step := drop
		for step > 0 && orders >= width {
			if orders >= step*(width) {
				orders -= width * step
				row := (height + height - step + 1) * step / 2
				res = (((row%mod)*(width%mod))%mod + res) % mod
				drop -= step
				height -= step
				step = drop
			} else {
				step /= 2
			}
		}
		if drop > 0 && orders < width {
			res = (((orders%mod)*(height%mod))%mod + res) % mod
			break
		}

		// 无优化
		// if orders >= len(inventory)-point {
		// 	orders -= len(inventory) - point
		// 	res = (((height%mod)*((len(inventory)-point)%mod))%mod + res) % mod
		// 	drop--
		// 	height--
		// } else {
		// 	res = (((orders%mod)*(height%mod))%mod + res) % mod
		// 	break
		// }
	}
	return res
}

func quickSort(A []int) {
	j := partition(A, 0, len(A))
	if j > 0 {
		quickSort(A[:j])
	}
	if j < len(A)-1 {
		quickSort(A[j+1:])
	}
}

func partition(A []int, lo, hi int) int {
	pivot := A[lo]
	i := lo + 1
	for j := lo + 1; j < hi; j++ {
		if A[j] < pivot {
			A[i], A[j] = A[j], A[i]
			i++
		}
	}
	A[lo], A[i-1] = A[i-1], A[lo]
	return i - 1
}
