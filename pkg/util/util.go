package util

import (
	"regexp"
	"unicode"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// validatePassword checks if the password is strong and meets the criteria:
// - At least 8 characters long
// - Contains at least one digit
// - Contains at least one lowercase letter
// - Contains at least one uppercase letter
// - Contains at least one special character
func ValidatePassword(pw string) (errMsg string, ok bool) {
	var (
		hasUpper, hasLower, hasNumber, hasSpecial bool
	)

	if len(pw) < 8 {
		return "Password must be at least 8 characters", false
	}

	for _, c := range pw {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsDigit(c):
			hasNumber = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}

	if !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return "Password must contain at least one of each category: uppercase letter, lowercase letter, number, and special character", false
	}

	return "", true
}
