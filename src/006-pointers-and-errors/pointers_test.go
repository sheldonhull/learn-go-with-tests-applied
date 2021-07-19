package pointers_test

import (
	"testing"

	"pointers"

	iz "github.com/matryer/is"
)

func TestWallet(t *testing.T) {
	t.Run("deposit a single amount", func(t *testing.T) {
		is := iz.New(t)

		wallet := pointers.Wallet{}
		wallet.Deposit(pointers.Bitcoin(10))

		got := wallet.Balance()
		want := pointers.Bitcoin(10)

		is.Equal(got, want) // balance returned
	})
	t.Run("deposit 2 amounts", func(t *testing.T) {
		is := iz.New(t)

		wallet := pointers.Wallet{}
		wallet.Deposit(pointers.Bitcoin(10))
		wallet.Deposit(pointers.Bitcoin(5))
		got := wallet.Balance()
		want := pointers.Bitcoin(15)

		is.Equal(got, want) // balance returned
	})

	t.Run("withdraw a single amount", func(t *testing.T) {
		is := iz.New(t)
		wallet := pointers.Wallet{}
		wallet.Deposit(pointers.Bitcoin(20))
		wallet.Withdraw(pointers.Bitcoin(10))
		got := wallet.Balance()
		want := pointers.Bitcoin(10)
		is.Equal(got, want) // withdraw returns reduced amount
	})
	t.Run("withdraw 2 amounts", func(t *testing.T) {
		is := iz.New(t)
		wallet := pointers.Wallet{}
		wallet.Deposit(pointers.Bitcoin(20))
		wallet.Withdraw(pointers.Bitcoin(5))
		got := wallet.Balance()
		want := pointers.Bitcoin(15)
		is.Equal(got, want) // withdraw returns reduced amount
	})
}
