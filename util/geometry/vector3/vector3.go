package vector3

import "math"

// Vector3 models a 3-dimensional vector.
type Vector3 struct {
	X, Y, Z float64
}

// Add adds a vector onto this vector.
func (v Vector3) Add(other Vector3) Vector3 {
	return Vector3{X: v.X + other.X, Y: v.Y + other.Y, Z: v.Z + other.Z}
}

// ScalarMult multiplies a scalar with a vector.
func (v Vector3) ScalarMult(scalar float64) Vector3 {
	return Vector3{X: scalar * v.X, Y: scalar * v.Y, Z: scalar * v.Z}
}

// Length returns the vector length
func (v Vector3) Length() float64 {
	return math.Sqrt(math.Pow(float64(v.X), 2) + math.Pow(float64(v.Y), 2) + math.Pow(float64(v.Z), 2))
}
