// this code copy in https://pdfs.semanticscholar.org/7679/bd13f5987da282b0662a0c41a4d6dd6f2165.pdf

package pi

func f(a float64) float64 {
	return 4.0 / (1.0 + a*a)
}

func chunk(start, end int64, c chan float64, h float64) {
	var sum float64 = 0.0
	for i := start; i < end; i++ {
		x := h * (float64(i) + 0.5)
		sum += f(x)
	}
	c <- sum * h
}

func chunknormal(start, end int64, h float64) float64 {
	var sum float64 = 0.0
	for i := start; i < end; i++ {
		x := h * (float64(i) + 0.5)
		sum += f(x)
	}
	return sum * h
}
