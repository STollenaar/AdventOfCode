package internal

import "math"

func abs(a int) int {
	return int(math.Abs(float64(a)))
}
