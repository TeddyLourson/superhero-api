package pkgtypes

import (
	"fmt"
	"regexp"
)

type Email struct {
	String string
}

func NewEmail(email string) (*Email, error) {
	if !ValidateEmail(email) {
		return nil, fmt.Errorf("invalid email: %s", email)
	}
	return &Email{String: email}, nil
}

func ValidateEmail(maybeEmail string) bool {
	isValidEmail, err := regexp.Match("(?i)"+"[A-Z0-9+_.-]+@[A-Z0-9.-]+\\.[A-Z]{2,}", []byte(maybeEmail))
	if err != nil {
		return false
	}
	return isValidEmail
}
