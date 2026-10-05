package main

import (
	"fmt"
	"log"
	"time"

	"example.com/finance-services/ledger/internal/ledger"
)

func main() {
	fmt.Println("Ledger service started")

	examples := []ledger.Transaction{
		{Amount: 1250.50, Category: "Продукты", Description: "Покупка продуктов", Date: time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC)},
		{Amount: 350, Category: "Транспорт", Description: "Поездка на такси", Date: time.Date(2026, 10, 2, 0, 0, 0, 0, time.UTC)},
		{Amount: 890, Category: "Развлечения", Description: "Билет в кино", Date: time.Date(2026, 10, 3, 0, 0, 0, 0, time.UTC)},
	}
	for _, tx := range examples {
		if err := ledger.AddTransaction(tx); err != nil {
			log.Fatalf("add transaction: %v", err)
		}
	}

	for _, tx := range ledger.ListTransactions() {
		fmt.Printf("ID=%d | Amount=%.2f | Category=%s | Description=%s | Date=%s\n",
			tx.ID, tx.Amount, tx.Category, tx.Description, tx.Date.Format("2006-01-02"))
	}
}
