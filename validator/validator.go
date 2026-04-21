package validator

import (
	"strings"
	"unicode"
)

func ValidateDomain(domain string) bool {
	if len(domain) == 0 {
		return false
	}
	authorized := "-."
	checkPeriod := false
	checkHyphen := false
	for i, ch := range domain {
		if i == 0 && strings.ContainsRune(authorized, ch) {
			return false
		}
		if !(isAlphaNum(ch) || strings.ContainsRune(authorized, ch)) {
			return false
		}
		if ch == '-' {
			// A period cannot sit next to an hyphen
			if checkPeriod == true {
				return false
			}
			checkHyphen = true
		}
		if ch == '.' {
			// An hyphen cannot sit next to an period
			if checkHyphen == true {
				return false
			}
			checkPeriod = true
		}
		if isAlphaNum(ch) && checkPeriod == true {
			checkPeriod = false
		}
		if isAlphaNum(ch) && checkHyphen == true {
			checkPeriod = false
		}
	}
	if checkPeriod == true || checkHyphen == true {
		return false
	}
	return true
}

func isAlphaNum(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch)
}
