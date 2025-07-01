package bank

import (
	"github.com/gnori-zon/go-tdd/generics/assert"
	"testing"
)

func TestBadBank(t *testing.T) {
	var (
		riya  = Account{Name: "Riya", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}

		transactions = []Transaction{
			NewTransaction(chris, riya, 100),
			NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(transactions, account).Balance
	}

	assert.Equal(t, newBalanceFor(riya), 200)
	assert.Equal(t, newBalanceFor(chris), 0)
	assert.Equal(t, newBalanceFor(adil), 175)
}
