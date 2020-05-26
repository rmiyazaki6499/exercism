package acronym

import (
	"strings"
)

func Abbreviate(s string) string {
	s = strings.Replace(s, "-", " ", -1)
	fmt.Println(s)
	words := strings.Fields(s)
	acro := ""
	for i := range words {
		acro += string(words[i][0])
	}
	return strings.ToUpper(acro)
}
