package problem1390

import (
	"math"
)

func sumFourDivisors(nums []int) int {
	sum := 0
	factors := make([]int, 0)
	for _, num := range nums {
		factors = factors[:0]
		square := int(math.Sqrt(float64(num)))
		for i := 1; i <= square; i++ {
			if num%i == 0 {
				if num == i*i {
					factors = append(factors, i)
					break
				} else {
					factors = append(factors, i, num/i)
				}
			}
		}
		if len(factors) == 4 {
			for i := 0; i < 4; i++ {
				sum += factors[i]
			}
		}
	}
	return sum
}
