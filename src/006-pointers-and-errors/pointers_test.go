package pointers_test

import (
	"testing"

	iz "github.com/matryer/is"
	"pointers"
)

func TestWallet(t *testing.T) {
	is := iz.New(t)
	assertBalance := func(t *testing.T, wallet *pointers.Wallet, want pointers.Bitcoin) {
		t.Helper()
		is := iz.New(t)
		got := wallet.Balance()
		is.Equal(got, want) // balance matches expected
	}
	t.Run("deposit a single amount", func(t *testing.T) {
		wallet := pointers.Wallet{}
		wallet.Deposit(pointers.Bitcoin(10))
		assertBalance(t, &wallet, pointers.Bitcoin(10))
	})
	t.Run("deposit 2 amounts", func(t *testing.T) {
		wallet := pointers.Wallet{}
		wallet.Deposit(pointers.Bitcoin(10))
		wallet.Deposit(pointers.Bitcoin(5))
		assertBalance(t, &wallet, pointers.Bitcoin(15))
	})

	t.Run("withdraw a single amount", func(t *testing.T) {
		wallet := pointers.Wallet{}
		wallet.Deposit(pointers.Bitcoin(20))
		err := wallet.Withdraw(pointers.Bitcoin(10))
		is.NoErr(err) // Withdraw should have no errors

		assertBalance(t, &wallet, pointers.Bitcoin(10))
	})
	t.Run("withdraw 2 amounts", func(t *testing.T) {
		wallet := pointers.Wallet{}
		wallet.Deposit(pointers.Bitcoin(20))
		err := wallet.Withdraw(pointers.Bitcoin(5))
		is.NoErr(err) // Withdraw should have no errors
		assertBalance(t, &wallet, pointers.Bitcoin(15))
	})

	t.Run("withdraw insufficient amount returns an error", func(t *testing.T) {
		wallet := pointers.Wallet{}
		wallet.Deposit(pointers.Bitcoin(20))
		err := wallet.Withdraw(pointers.Bitcoin(100))
		is.True(err != nil) // wanted an error but didn't get one
	})

	t.Run("balance is returned", func(t *testing.T) {
		wallet := pointers.Wallet{}
		wallet.Deposit(pointers.Bitcoin(20))
		got := wallet.Balance()
		want := pointers.Bitcoin(20)
		is.Equal(got, want) // balance returns what it gets
	})
}
