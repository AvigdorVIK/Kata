package kata

func PositiveSum(numbers []int) int {
	unswer := 0
	for _, s := range numbers {
		if s > 0 {
			unswer += s
		}

	}
	return unswer
}
