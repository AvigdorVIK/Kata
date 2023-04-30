package kata

import "strings"

func FindShort(s string) int {
	word := strings.Split(s, " ")
	shortWord := len(word[0])
	for _, wort := range word {
		if len(wort) < shortWord {
			shortWord = len(wort)
		}
	}
	return shortWord
}
