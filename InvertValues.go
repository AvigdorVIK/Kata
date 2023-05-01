package kata

func Invert(arr []int) []int {
	arr1 := make([]int, len(arr))

	for i, nummer := range arr {
		nummer = -nummer
		arr1[i] = nummer
	}
	return arr1
}
