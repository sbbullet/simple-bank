package db

import (
	"context"
	"testing"
	"time"

	"github.com/sbbullet/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomBalance(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.FromAccountID, arg.FromAccountID)
	require.Equal(t, transfer.ToAccountID, arg.ToAccountID)
	require.Equal(t, transfer.Amount, arg.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	var transfer Transfer

	for i := 0; i < 10; i++ {
		transfer = createRandomTransfer(t)
	}

	arg := ListTransfersParams{
		Limit:         5,
		Offset:        0,
		FromAccountID: transfer.FromAccountID,
		ToAccountID:   transfer.ToAccountID,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(transfers), 1)
	require.LessOrEqual(t, len(transfers), 5)

	for _, singleTransfer := range transfers {
		require.NotEmpty(t, singleTransfer)
	}
}
