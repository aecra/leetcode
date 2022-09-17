package lcp55

import "math"

func getMinimumTime(time []int, fruits [][]int, limit int) int {
	minTime := 0
	for _, fruit := range fruits {
		minTime += time[fruit[0]] * int(math.Ceil(float64(fruit[1])/float64(limit)))
	}
	return minTime
}
