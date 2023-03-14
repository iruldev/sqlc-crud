// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package database

import (
	"context"
)

type Querier interface {
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) error
	CreateAccount(ctx context.Context, arg CreateAccountParams) (int64, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (int64, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (int64, error)
	CreateUser(ctx context.Context, arg CreateUserParams) error
	DeleteAccount(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetAccountByOwner(ctx context.Context, owner string) (Account, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListAccount(ctx context.Context, arg ListAccountParams) ([]Account, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) error
}

var _ Querier = (*Queries)(nil)
