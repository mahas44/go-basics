package test

func MySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func MyAverage(xi []float64) float64 {
	var sum float64 = 0
	for _, v := range xi {
		sum += v
	}
	return sum / float64(len(xi))
}
