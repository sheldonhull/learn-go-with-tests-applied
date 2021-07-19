package pointers

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(n int) {
	w.balance = n
	fmt.Printf("balance address: %v", &w.balance)
	s.Show("balance", &w.balance)
}

func (w *Wallet) Balance() int {
	return w.balance
}
