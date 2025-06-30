package bank

import "github.com/gnori-zon/go-tdd/slices"

type Account struct {
	Name    string
	Balance float64
}

type Transaction struct {
	From, To string
	Sum      float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

func NewBalanceFor(transactions []Transaction, account Account) Account {
	return slices.Reduce(transactions, applyTransaction, account)
}

func applyTransaction(account Account, transaction Transaction) Account {
	if transaction.From == account.Name {
		account.Balance -= transaction.Sum
	}
	if transaction.To == account.Name {
		account.Balance += transaction.Sum
	}
	return account
}
