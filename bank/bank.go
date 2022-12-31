package bank

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(bitcoins Bitcoin) {
	w.balance += bitcoins
}

func (w *Wallet) Withdraw(bitcoins Bitcoin) error {
	if bitcoins > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= bitcoins
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
