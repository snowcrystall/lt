package medium

import (
	"math"
	"math/rand"
)

type Solution struct {
	radius, xCenter, yCenter float64
}

func Constructor(radius, xCenter, yCenter float64) Solution {
	return Solution{radius, xCenter, yCenter}
}

func (s *Solution) RandPoint() []float64 {
	r := math.Sqrt(rand.Float64())
	sin, cos := math.Sincos(rand.Float64() * 2 * math.Pi)
	return []float64{s.xCenter + r*cos*s.radius, s.yCenter + r*sin*s.radius}
}
