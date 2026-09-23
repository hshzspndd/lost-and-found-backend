package utils

import "regexp"

var phoneRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)

func IsValidPhone(phone string) bool {
	return phoneRegex.MatchString(phone)
}
