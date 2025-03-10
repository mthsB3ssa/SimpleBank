// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID        int32            `json:"id"`
	Owner     string           `json:"owner"`
	Balance   pgtype.Numeric   `json:"balance"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type Entry struct {
	ID        int32            `json:"id"`
	AccountID int32            `json:"account_id"`
	Amount    pgtype.Numeric   `json:"amount"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type Transfer struct {
	ID            int32            `json:"id"`
	FromAccountID int32            `json:"from_account_id"`
	ToAccountID   int32            `json:"to_account_id"`
	Amount        pgtype.Numeric   `json:"amount"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
}
