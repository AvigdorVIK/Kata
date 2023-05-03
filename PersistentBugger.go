package kata

func Persistence(n int) int {
	count := 0
	for n >= 10 {
		number := 1
		for n > 0 {
			number *= n % 10
			n = n / 10
		}
		n = number
		count += 1
	}
	return count
}
