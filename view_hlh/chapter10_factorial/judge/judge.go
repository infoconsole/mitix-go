package judge

func MultiplySteps(x int, piper chan int) {
	if x == 0 {
		piper <- 1
		return
	}
	var ret = 1
	for i := 1; i <= x; i++ {
		ret = ret * i
	}
	piper <- ret
	return
}
