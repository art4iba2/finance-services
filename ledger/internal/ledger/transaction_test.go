package ledger

import (
	"reflect"
	"testing"
	"time"
)

func TestTransactionStore(t *testing.T) {
	transactions = make([]Transaction, 0)
	t.Cleanup(func() { transactions = make([]Transaction, 0) })

	if got := ListTransactions(); got == nil || len(got) != 0 {
		t.Fatalf("new store should return an empty slice, got %v", got)
	}

	first := Transaction{ID: 99, Amount: 100, Category: "Food", Description: "Lunch", Date: time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC)}
	if err := AddTransaction(first); err != nil {
		t.Fatal(err)
	}
	if err := AddTransaction(Transaction{Amount: 0}); err == nil {
		t.Fatal("zero amount must be rejected")
	}
	second := Transaction{Amount: 50, Category: "Transport"}
	if err := AddTransaction(second); err != nil {
		t.Fatal(err)
	}

	first.ID, second.ID = 1, 2
	want := []Transaction{first, second}
	got := ListTransactions()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
	}

	got[0].Amount = 999
	if !reflect.DeepEqual(ListTransactions(), want) {
		t.Fatal("changing the returned slice must not change stored transactions")
	}
}
