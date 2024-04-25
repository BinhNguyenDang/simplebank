package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullname = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("invalid length: must be between %d and %d characters", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("invalid characters: must be lowercase alphanumeric, digits or underscore")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil{
		return fmt.Errorf("invalid email address: %s", err)
	}
	return nil
}

func ValidateFullName(value string) error {
    if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidFullname(value) {
		return fmt.Errorf("invalid characters: must be alphanumeric or spaces")
	}
	return nil
}

func ValidateEmailId(value int64) error {
	if value <= 0 {
        return fmt.Errorf("invalid email id: must be positive integer")
    }
    return nil
}

func ValidateSecretCode( value string ) error {
	return ValidateString(value, 32, 128)
}