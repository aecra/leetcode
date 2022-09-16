package problem935

func knightDialer(n int) int {
	// b1: 1、3、7、9
	// b2: 2、8
	// b3: 4、6
	// b4: 0
	// b5: 5
	b1, b2, b3, b4, b5 := 4, 2, 2, 1, 1
	for i := 1; i < n; i++ {
		b1, b2, b3, b4, b5 = (2*b2+2*b3)%1000000007, b1, (2*b4+b1)%1000000007, b3, 0
	}
	return (b1 + b2 + b3 + b4 + b5) % 1000000007
}
