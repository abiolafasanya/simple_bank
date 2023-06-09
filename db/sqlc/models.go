// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"database/sql"
	"time"
)

type Account struct {
	ID          int64         `db:"id"`
	Owner       string        `db:"owner"`
	Balance     int64         `db:"balance"`
	Currency    string        `db:"currency"`
	CountryCode sql.NullInt32 `db:"country_code"`
	CreatedAt   time.Time     `db:"created_at"`
}

type Entry struct {
	ID        int64 `db:"id"`
	AccountID int64 `db:"account_id"`
	// It can be negative or positive
	Account   int64     `db:"account"`
	CreatedAt time.Time `db:"created_at"`
}

type Transfer struct {
	ID            int64 `db:"id"`
	FromAccountID int64 `db:"from_account_id"`
	ToAccountID   int64 `db:"to_account_id"`
	// must be positive
	Account   int64     `db:"account"`
	CreatedAt time.Time `db:"created_at"`
}
