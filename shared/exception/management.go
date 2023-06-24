package exception

import (
	"errors"
)

func DrugNotFound() error {
	return errors.New("Drug not found")
}

func DrugAlreadyExist() error {
	return errors.New("Drug already exist")
}