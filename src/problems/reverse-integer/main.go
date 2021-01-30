package main

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	isNegative := false
	if x < 0 {
		isNegative = true
		x = x * -1
	}

	s := ""
	conS := strconv.Itoa(x)
	le := len(conS)
	for i := le - 1; i >= 0; i-- {
		s += string(conS[i])
	}

	n, _ := strconv.Atoi(s)

	if isNegative {
		n = n * -1
	}
	if n < math.MinInt32 || n > math.MaxInt32 {
		return 0
	}

	return n
}
func main() {
	x := -1
	reverse(x)
}
