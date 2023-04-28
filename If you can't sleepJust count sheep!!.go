package kata

import (
	"fmt"
	"strings"
)

func countSheep(num int) string {
	stri := []string{}
	for i := 1; i <= num; i++ {
		stri = append(stri, fmt.Sprintf("%d sheep...", i))
	}
	sentence := strings.Join(stri, "")
	return sentence
}
