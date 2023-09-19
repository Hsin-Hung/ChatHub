package utils

import (
	"errors"
	"strings"
)

func UserInputValidator(username string, password string) error {

	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		return errors.New("username and password cannot be empty")
	}

	if len(password) < 10 || len(password) > 20 {
		return errors.New("password length must be between 10 to 20 characters")
	}

	return nil
}
