package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/shivshankarm/bankservice/util"
	"github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Accounts {
	user := CreateRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.ID)
	return account
}

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account2.Owner, account1.Owner)
	require.Equal(t, account2.Balance, account1.Balance)
	require.Equal(t, account2.Currency, account1.Currency)
	require.WithinDuration(t, account2.CreatedAt, account1.CreatedAt, time.Second)
	require.Equal(t, account2.ID, account1.ID)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func TestUpdateAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}
	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account2.Owner, account1.Owner)
	require.Equal(t, account2.Balance, arg.Balance)
	require.Equal(t, account2.Currency, account1.Currency)
	require.WithinDuration(t, account2.CreatedAt, account1.CreatedAt, time.Second)
	require.Equal(t, account2.ID, account1.ID)
}

func TestDeleteAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func CreateEntry(t *testing.T) Entries {
	account1 := CreateRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account1.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.Amount, entry.Amount)
	return entry
}
func TestCreateEntry(t *testing.T) {
	CreateEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := CreateEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry2.ID, entry1.ID)
	require.Equal(t, entry2.AccountID, entry1.AccountID)
	require.Equal(t, entry2.Amount, entry1.Amount)
}

func TestUpdateEntry(t *testing.T) {
	entry1 := CreateEntry(t)
	arg := UpdateEntryParams{
		AccountID: entry1.AccountID,
		Amount:    util.RandomMoney(),
	}
	entry2, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry2.Amount, arg.Amount)
}

func TestDeleteEntry(t *testing.T) {
	entry1 := CreateEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}
func TestListEntry(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateEntry(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

func CreateTransfer(t *testing.T) Transfers {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, arg.Amount, transfer.Amount)
	return transfer
}
func TestCreateTransfer(t *testing.T) {
	CreateTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := CreateTransfer(t)
	arg := GetTransferParams{
		FromAccountID: transfer1.FromAccountID,
		ToAccountID:   transfer1.ToAccountID,
	}
	transfer2, err := testQueries.GetTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer2.ID, transfer1.ID)
	require.Equal(t, transfer2.Amount, transfer1.Amount)
}
func TestUpdateTransfer(t *testing.T) {
	transfer1 := CreateTransfer(t)

	arg := UpdateTransferParams{
		FromAccountID: transfer1.FromAccountID,
		ToAccountID:   transfer1.ToAccountID,
		Amount:        util.RandomMoney(),
	}
	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer2.ID, transfer1.ID)
	require.Equal(t, transfer2.Amount, arg.Amount)
}

func TestDeleteTransfer(t *testing.T) {
	transfer1 := CreateTransfer(t)
	arg := GetTransferParams{
		FromAccountID: transfer1.FromAccountID,
		ToAccountID:   transfer1.ToAccountID,
	}
	err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	transfer2, err := testQueries.GetTransfer(context.Background(), arg)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transfer2)
}
func TestListTransfer(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateTransfer(t)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
