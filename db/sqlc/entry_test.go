package db

import (
	"context"
	"testing"
	"time"

	"github.com/sbbullet/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) (Account, Entry) {
	account := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomBalance(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, account.ID)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return account, entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	account, entry1 := createRandomEntry(t)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, account.ID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {

	var account Account

	for i := 0; i < 10; i++ {
		account, _ = createRandomEntry(t)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    0,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(entries), 1)
	require.LessOrEqual(t, len(entries), 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
