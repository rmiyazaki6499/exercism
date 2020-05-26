package bob

import (
	"strings"
	"unicode"
)

type Remark struct {
	remark string
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func isShouting(remark string) bool {
	isUppercased := strings.ToUpper(remark) == remark
	hasLetters := strings.IndexFunc(remark, unicode.IsLetter) >= 0
	return isUppercased && hasLetters
}

func isQuiet(remark string) bool {
	return remark == ""
}

func isShoutingQuestion(remark string) bool {
	return isQuestion(remark) && isShouting(remark)
}

func Hey(remark string) string {
	newRemark := strings.TrimSpace(remark)
	switch {
	case isShoutingQuestion(newRemark):
		return "Calm down, I know what I'm doing!"
	case isQuestion(newRemark):
		return "Sure."
	case isShouting(newRemark):
		return "Whoa, chill out!"
	case isQuiet(newRemark):
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
