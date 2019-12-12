package calc

import "math"

// GreatestCommonDivisor returns the gcd of a and b
// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GreatestCommonDivisor(a, b int64) int64 {
	a = int64(math.Abs(float64(a)))
	b = int64(math.Abs(float64(b)))

	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LeastCommonMultiple returns the lcm of all passed integers
// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func LeastCommonMultiple(a, b int64, integers ...int64) int64 {
	result := a * b / GreatestCommonDivisor(a, b)

	for i := 0; i < len(integers); i++ {
		result = LeastCommonMultiple(result, integers[i])
	}

	return result
}
