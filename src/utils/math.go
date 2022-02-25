package utils

import "math"

func Round2(value float64) float64 {
	return math.Round(value*100) / 100
}
