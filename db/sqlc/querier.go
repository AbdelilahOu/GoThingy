// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"context"
)

type Querier interface {
	GetAcount(ctx context.Context, id int64) (Account, error)
}

var _ Querier = (*Queries)(nil)
