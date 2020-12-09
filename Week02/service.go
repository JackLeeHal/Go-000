package Week02

import "errors"

func CommentsService(id string) ([]Comment, error) {
	c, err := GetComments(id)
	if err != nil {
		// log
		var e *CustomError
		if errors.As(err, &e) {
			return []Comment{}, nil
		}

		return nil, err
	}

	return c, nil
}

func BalanceService(id string) (*Balance, error) {
	b, err := GetBalance(id)
	if err != nil {
		// log
		return nil, err
	}

	return b, nil
}
