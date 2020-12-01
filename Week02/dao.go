package Week02

import "errors"

var (
	errCommentsNotFound = errors.New("comments not found")
	errBalanceNotFound  = errors.New("balance not found")
)

func GetComments(id string, db map[string][]Comment) ([]Comment, error) {
	comments := db[id]
	if len(comments) == 0 {
		return nil, errCommentsNotFound
	}
	return comments, nil
}

func GetBalance(id string, db map[string]*Balance) (*Balance, error) {
	balance := db[id]
	if balance == nil {
		return nil, errBalanceNotFound
	}
	return balance, nil
}
