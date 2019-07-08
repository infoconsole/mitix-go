package judge

import "strconv"

func DataJudge(x int, piper chan int, pipet chan int) {
	xstr := strconv.Itoa(x)
	var actualx int
	for i := 0; i < len(xstr); i++ {
		var xi int
		xi, _ = strconv.Atoi(xstr[i : i+1])
		actualx = actualx + (xi * xi * xi)
	}
	pipet <- x
	if x == actualx {
		piper <- x
	}
	return
}
