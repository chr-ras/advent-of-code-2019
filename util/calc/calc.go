package calc

import "math"

// GreatestCommonDivisor returns the gcd of a and b
func GreatestCommonDivisor(a, b int) int {
	a = int(math.Abs(float64(a)))
	b = int(math.Abs(float64(b)))

	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
