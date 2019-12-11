package geometry

import (
	"math"

	"github.com/chr-ras/advent-of-code-2019/util/calc"
)

// Point represents a point in 2D space
type Point struct {
	X, Y int
}

// Vector represents a vector in 2D space
type Vector struct {
	X, Y int
}

// GetDirectionVector determines the vector needed to get from one point to another point
func GetDirectionVector(from, to Point) Vector {
	return Vector{X: to.X - from.X, Y: to.Y - from.Y}
}

// GetSmallestFactorVector determines the smallest factor vector in the given vector.
// Examples: {1, 1} => {1, 1}; {2, 2} => {1, 1}; {2, 4} => {1, 2}; {3, 5} => {3, 5}
func (v Vector) GetSmallestFactorVector() Vector {
	gcd := calc.GreatestCommonDivisor(v.X, v.Y)

	return Vector{X: v.X / gcd, Y: v.Y / gcd}
}

// ScalarMult multiplies a scalar with a vector.
func (v Vector) ScalarMult(scalar int) Vector {
	return Vector{X: v.X * scalar, Y: v.Y * scalar}
}

// AsPoint transforms a vector into a point.
func (v Vector) AsPoint() Point {
	return Point{X: v.X, Y: v.Y}
}

// Add returns the result of the vector addition.
func (v Vector) Add(other Vector) Vector {
	return Vector{X: v.X + other.X, Y: v.Y + other.Y}
}

// Length returns the vector length
func (v Vector) Length() float64 {
	return math.Sqrt(math.Pow(float64(v.X), 2) + math.Pow(float64(v.Y), 2))
}

// AsVector transforms a point into a vector.
func (p Point) AsVector() Vector {
	return Vector{X: p.X, Y: p.Y}
}

// GetAngle produces the angle on the side of p with the other point
func (p Point) GetAngle(other Point) float64 {
	rads := math.Atan2(float64(other.Y-p.Y), float64(other.X-p.X))
	rads = math.Mod(rads+2*math.Pi, 2*math.Pi)

	return math.Mod(rads+math.Pi/2, 2*math.Pi)
}
