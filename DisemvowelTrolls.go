package kata

import "strings"

func Disemvowel(comment string) string {
	vowels := "aeiouAEIOU"
	answer := ""

	for _, cha := range comment {
		if !strings.ContainsRune(vowels, cha) {
			answer += string(cha)
		}
	}
	return answer
}
