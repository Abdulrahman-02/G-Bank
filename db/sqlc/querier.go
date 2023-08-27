// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"context"
)

type Querier interface {
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Accounts, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Accounts, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entries, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfers, error)
	DeleteAccount(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Accounts, error)
	GetAccountForUpdate(ctx context.Context, id int64) (Accounts, error)
	GetEntry(ctx context.Context, id int64) (Entries, error)
	GetTransfer(ctx context.Context, id int64) (Transfers, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Accounts, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entries, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfers, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Accounts, error)
}

var _ Querier = (*Queries)(nil)
