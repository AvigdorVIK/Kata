package kata

func Gps(s int, x []float64) int {
	if len(x) <= 1 {
		return 0 // Return 0 if the car didn't move
	}

	maxAvgSpeed := 0
	for i := 0; i < len(x)-1; i++ {
		deltaDistance := x[i+1] - x[i]
		avgSpeed := int((3600 * deltaDistance) / float64(s))

		if avgSpeed > maxAvgSpeed {
			maxAvgSpeed = avgSpeed
		}
	}

	return maxAvgSpeed
}
