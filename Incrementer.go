package kata

func Incrementer(n []int) []int {
	result := make([]int, len(n))
	for i, num := range n {
		incrementedNum := num + (i + 1)
		result[i] = incrementedNum % 10
	}
	return result
}
