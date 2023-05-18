package kata

import "strings"

func GetCount(str string) int {
	count := 0
	strs := strings.ToLower(str)

	for _, vowel := range strs {
		if vowel == 'a' || vowel == 'e' || vowel == 'i' || vowel == 'o' || vowel == 'u' {
			count++
		}
	}
	return count
}
