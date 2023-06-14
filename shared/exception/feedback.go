package exception

import (
	"errors"
)

func FeedbackCategoryAlreadyExist() error {
	return errors.New("Feedback category already exist")
}