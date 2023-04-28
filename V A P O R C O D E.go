package kata

import "strings"

func Vaporcode(s string) string {
	var builder strings.Builder
	s = strings.ToUpper(s)
	for _, char := range s {
		if char == ' ' {
			builder.WriteString("")
		} else {
			builder.WriteRune(char)
			builder.WriteString("  ")
		}
	}
	return strings.TrimSpace(builder.String())

}
