package db

import (
	"context"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:   "John",
		Balance: 1000,
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	if err != nil {
		t.Fatal("failed to create account:", err)
	}
	if account.Owner != arg.Owner {
		t.Errorf("expected owner %s, got %s", arg.Owner, account.Owner)
	}
}
