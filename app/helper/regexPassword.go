package helper

import (
	"unicode"
)

func ValidatePassword(password string) bool {
	// Check password length
	if len(password) < 8 {
		return false
	}

	// Check if password contains at least one uppercase letter, one symbol, and one alphanumeric character
	var (
		hasUppercase  bool
		hasSymbol     bool
		hasAlpha 			bool
		hasNumeric		bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSymbol = true
		case unicode.IsDigit(char):
			hasNumeric = true
		case unicode.IsLetter(char):
			hasAlpha = true
		}
	}

	return hasUppercase && hasSymbol && hasAlpha && hasNumeric
}