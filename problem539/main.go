package problem539

import "sort"

func findMinDifference(timePoints []string) int {
	times := make([]int, 0, len(timePoints)+1)
	for _, timePoint := range timePoints {
		minute := 0
		minute += (int(timePoint[0]-'0')*10 + int(timePoint[1]-'0')) * 60
		minute += (int(timePoint[3]-'0')*10 + int(timePoint[4]-'0'))
		times = append(times, minute)
	}
	sort.Ints(times)
	times = append(times, times[0]+24*60)
	min := times[1] - times[0]
	for i := 1; i < len(times); i++ {
		if times[i]-times[i-1] < min {
			min = times[i] - times[i-1]
		}
	}
	return min
}
