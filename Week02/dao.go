package Week02

import (
	"errors"
	"fmt"
)

var (
	ErrNoRows  = errors.New("no rows found")
	commentsDB = map[string][]Comment{
		"test": {},
	}
	balanceDB = map[string]*Balance{}
)

type CustomError struct {
	name  string
	cause error
}

func (e *CustomError) Error() string {
	return fmt.Sprint("error ", e.name)
}

func GetComments(id string) ([]Comment, error) {
	comments := commentsDB[id]
	if len(comments) == 0 {
		return nil, &CustomError{
			name:  "Comments",
			cause: ErrNoRows,
		}
	}
	return comments, nil
}

func GetBalance(id string) (*Balance, error) {
	balance := balanceDB[id]
	if balance == nil {
		return nil, &CustomError{
			name:  "Balance",
			cause: ErrNoRows,
		}
	}
	return balance, nil
}
