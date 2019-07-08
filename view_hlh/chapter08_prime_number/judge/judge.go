package judge

import "math"

func PrimeJupdge(x int, piper chan int, pipet chan int) {
	xx := int(math.Abs(float64(x)))
	//其实不需要到xx  只需要到xx的平方根就好了
	for i := 2; i < int(math.Sqrt(float64(xx))); i++ {
		if xx%i == 0 {
			pipet <- x
			return
		}
	}
	pipet <- x
	piper <- x
	return
}
