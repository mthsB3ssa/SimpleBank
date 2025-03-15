package db

import "testing"

func TestCreateAccount(t *testing.T){
	arg := CreateAccountParams{
		Owner: "John",
		Balance: 1000,
	}
}
