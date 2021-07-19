package pointers_test

import (
	"testing"

	iz "github.com/matryer/is"

	"pointers"
)

func TestWallet(t *testing.T) {
	is := iz.New(t)

	wallet := pointers.Wallet{}
	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	is.Equal(got, want) // balance returned
}
