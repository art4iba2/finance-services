package ledger

import (
	"errors"
	"time"
)

// Transaction represents a financial transaction.
type Transaction struct {
	ID          int
	Amount      float64
	Category    string
	Description string
	Date        time.Time
}

// This in-memory store is initialized when the package is loaded.
// The demo accesses it sequentially; it is not safe for concurrent use.
var transactions = make([]Transaction, 0)

// AddTransaction assigns an ID and saves a non-zero transaction.
func AddTransaction(tx Transaction) error {
	if tx.Amount == 0 {
		return errors.New("transaction amount must not be zero")
	}

	tx.ID = len(transactions) + 1
	transactions = append(transactions, tx)
	return nil
}

// ListTransactions returns a copy so callers cannot modify stored values.
func ListTransactions() []Transaction {
	result := make([]Transaction, len(transactions))
	copy(result, transactions)
	return result
}
