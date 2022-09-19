package problem1013

func canThreePartsEqualSum(arr []int) bool {
	prefix := make([]int, len(arr))
	prefix[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		prefix[i] += prefix[i-1] + arr[i]
	}

	if prefix[len(arr)-1]%3 != 0 {
		return false
	}

	everyPart := prefix[len(prefix)-1] / 3
	for i := 0; i < len(prefix)-2; i++ {
		if prefix[i] == everyPart {
			for j := i + 1; j < len(prefix)-1; j++ {
				if prefix[j]-prefix[i] == everyPart {
					return true
				}
			}
		}
	}
	return false
}
