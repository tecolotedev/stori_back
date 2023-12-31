// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: account.sql

package sqlc_code

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
    balance,
    currency,
    user_id
)
VALUES (
    $1, $2, $3
) RETURNING id, balance, currency, created_at, user_id
`

type CreateAccountParams struct {
	Balance  pgtype.Float8
	Currency string
	UserID   pgtype.Int4
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, createAccount, arg.Balance, arg.Currency, arg.UserID)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
		&i.UserID,
	)
	return i, err
}

const getAccount = `-- name: GetAccount :one
SELECT id, balance, currency, created_at, user_id FROM accounts
where id =$1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int32) (Account, error) {
	row := q.db.QueryRow(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
		&i.UserID,
	)
	return i, err
}

const getAccountForUpdate = `-- name: GetAccountForUpdate :one
SELECT id, balance, currency, created_at, user_id FROM accounts
where id =$1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetAccountForUpdate(ctx context.Context, id int32) (Account, error) {
	row := q.db.QueryRow(ctx, getAccountForUpdate, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
		&i.UserID,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, balance, currency, created_at, user_id FROM accounts
WHERE user_id = $1
ORDER BY id 
LIMIT $2
OFFSET $3
`

type ListAccountsParams struct {
	UserID pgtype.Int4
	Limit  int32
	Offset int32
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	rows, err := q.db.Query(ctx, listAccounts, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Balance,
			&i.Currency,
			&i.CreatedAt,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE accounts
SET balance =$2
WHERE id = $1
RETURNING id, balance, currency, created_at, user_id
`

type UpdateAccountParams struct {
	ID      int32
	Balance pgtype.Float8
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, updateAccount, arg.ID, arg.Balance)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
		&i.UserID,
	)
	return i, err
}
