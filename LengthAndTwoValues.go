package kata

func Alternate(n int, firstValue string, secondValue string) (slice []string) {
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			slice = append(slice, firstValue)
		} else {
			slice = append(slice, secondValue)
		}
	}
	return slice
}
