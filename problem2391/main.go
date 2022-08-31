package problem2391

func garbageCollection(garbage []string, travel []int) int {
	metalRubbish := make([]int, len(garbage))
	paperRubbish := make([]int, len(garbage))
	glassRubbish := make([]int, len(garbage))
	for i, g := range garbage {
		for _, r := range g {
			switch r {
			case 'M':
				metalRubbish[i]++
			case 'P':
				paperRubbish[i]++
			case 'G':
				glassRubbish[i]++
			}
		}
	}

	// count metal cost
	metalCost := metalRubbish[0]
	metalEnd := endAt(metalRubbish)
	for i := 1; i <= metalEnd; i++ {
		metalCost += metalRubbish[i] + travel[i-1]
	}
	// count paper cost
	paperCost := paperRubbish[0]
	paperEnd := endAt(paperRubbish)
	for i := 1; i <= paperEnd; i++ {
		paperCost += paperRubbish[i] + travel[i-1]
	}
	// count glass cost
	glassCost := glassRubbish[0]
	glassEnd := endAt(glassRubbish)
	for i := 1; i <= glassEnd; i++ {
		glassCost += glassRubbish[i] + travel[i-1]
	}

	return metalCost + paperCost + glassCost
}

func endAt(rubbish []int) int {
	for i := len(rubbish) - 1; i >= 0; i-- {
		if rubbish[i] > 0 {
			return i
		}
	}
	return -1
}
