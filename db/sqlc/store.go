package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()

}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfers `json:"transfer"`
	FromAccount Accounts  `json:"from_account"`
	ToAccount   Accounts  `json:"to_account"`
	FromEntry   Entries   `json:"from_entry"`
	ToEntry     Entries   `json:"to_entry"`
}

func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			arg.FromAccountID,
			arg.ToAccountID,
			arg.Amount,
		})
		if err != nil {
			return nil
		}
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			arg.FromAccountID,
			-arg.Amount,
		})
		if err != nil {
			return nil
		}
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			arg.ToAccountID,
			arg.Amount,
		})
		if err != nil {
			return nil
		}

		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, err = addAcountBalance(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, err = addAcountBalance(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
		}

		return nil
	})

	return result, err
}

func addAcountBalance(
	ctx context.Context,
	q *Queries,
	account1_id int64,
	amount1 int64,
	account2_id int64,
	amount2 int64,
) (account1 Accounts, account2 Accounts, err error) {
	// move money out of account1
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     account1_id,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	// move money into account2
	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     account2_id,
		Amount: amount2,
	})
	return
}
