package pointersanderrors

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type Bitcoin int64

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(value Bitcoin) {
	w.balance += value
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(value Bitcoin) error {
	if w.balance < value {
		return ErrInsufficientFunds
	}
	w.balance -= value
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
