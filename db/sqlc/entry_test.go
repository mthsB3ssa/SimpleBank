package db

import (
	"context"
	"testing"

	"github.com/mthsB3ssa/SimpleBank/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T) Entry {
	arg := CreateEntryParams{
		AccountID: utils,
		Amount:    utils.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
}

func TestCreateEntry(t *testing.T) {
	account := CreateRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount: utils.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.NotEmpty(t, arg.Amount)
}
func TestGetEntry(t *testing.T) {

}

func TestListEntries(t *testing.T) {

}
