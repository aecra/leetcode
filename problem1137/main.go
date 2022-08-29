package problem1137

var tribonacciNum = map[int]int{
	0: 0,
	1: 1,
	2: 1,
}

func tribonacci(n int) int {
	if num, ok := tribonacciNum[n]; ok {
		return num
	}
	num := tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
	tribonacciNum[n] = num
	return num
}
