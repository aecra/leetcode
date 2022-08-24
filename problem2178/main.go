package problem2178

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 == 1 {
		return []int64{}
	}
	var res []int64
	var base int64 = 2
	for {
		if base > finalSum {
			res[len(res)-1] += finalSum
			break
		} else {
			res = append(res, base)
			finalSum -= base
			base += 2
		}
	}
	return res
}
