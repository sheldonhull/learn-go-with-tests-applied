/*
Package pointers manages a bitcoin wallet with magical unicorn dust.

Use this with care, and ignore the fact this is not using floats.
*/
package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin provides an integer value representing the Bitcoin value.
type Bitcoin int

// Stringer returns the string value
type Stringer interface {
	String() string
}

// Wallet contains the bitcoin balance
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Withdraw subtracts the provided bitcoin amount from the balance.
// If the bitcoin amount is insufficient then an error is returned.
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if w.balance-amount < 0 {
		return errors.New("insufficient balance")
	}
	w.balance -= amount

	return nil
}
