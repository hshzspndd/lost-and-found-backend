package utils

import (
	"regexp"
)

var phoneRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)

func IsValidPhone(phone string) bool {
	if !phoneRegex.MatchString(phone) {
		return false
	}
	return true
}
