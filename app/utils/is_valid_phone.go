package utils

import (
	"lost-and-found-backend/app/errs"
	"regexp"
)

var phoneRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)

func IsValidPhone(phone string) error {
	if !phoneRegex.MatchString(phone) {
		return errs.ErrPhoneFormat
	}
	return nil
}
