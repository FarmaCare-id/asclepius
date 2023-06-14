package exception

import (
	"errors"
)

func UserAlreadyExist() error {
	return errors.New("User already exist")
}

func UserNotFound() error {
	return errors.New("No user found for given email")
}

func IncorrectPassword() error {
	return errors.New("Incorrect password")
}

func InvalidToken() error {
	return errors.New("Invalid token")
}

func ExpiredToken() error {
	return errors.New("Expired token")
}