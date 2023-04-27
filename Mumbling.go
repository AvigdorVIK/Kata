package kata

import "strings"

func Accum(s string) string {
	var result []string
	for i, c := range s {
		result = append(result, strings.ToUpper(string(c))+strings.Repeat(strings.ToLower(string(c)), i))
	}
	return strings.Join(result, "-")
}
