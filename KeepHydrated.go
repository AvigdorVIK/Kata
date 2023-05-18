package kata

import "math"

func Litres(time float64) int {
	litres := math.Floor(time * 0.5)
	return int(litres)
}
