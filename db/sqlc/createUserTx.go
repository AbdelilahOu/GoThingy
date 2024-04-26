package db

import (
	"context"
)

// transfer tx params
type CreateUserTxParams struct {
	CreateUserParams
	AfterCreate func(user User) error
}

// transfer result
type CreateUserTxResult struct {
	User User
}

// specific transation
func (store *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.User, err = q.CreateUser(ctx, arg.CreateUserParams)
		if err != nil {
			return err
		}
		return arg.AfterCreate(result.User)
	})
	return result, err
}
