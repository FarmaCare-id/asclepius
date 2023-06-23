package exception

import (
	"errors"
)

func DrugAlreadyExist() error {
	return errors.New("Drug already exist")
}