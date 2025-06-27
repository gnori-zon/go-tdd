package bank

type Transaction struct {
	From, To string
	Sum      float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	balance := 0.0
	for _, transaction := range transactions {
		if transaction.From == name {
			balance -= transaction.Sum
		}
		if transaction.To == name {
			balance += transaction.Sum
		}
	}
	return balance
}
