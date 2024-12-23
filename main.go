package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Transaction struct {
	ID       int
	Amount   float64
	Category string
	Date     time.Time
	Type     string
}

type BudgetTracker struct {
	Transactions []Transaction
	nextID       int
}

type FinancialRecord interface {
	GetAmount() float64
	GetType() string
}

func (t *Transaction) GetAmount() float64 {
	return t.Amount
}

func (t *Transaction) GetType() string {
	return t.Type
}

func (bt *BudgetTracker) AddTransaction(amount float64, category, transactionType string) {
	newTransaction := Transaction{
		ID:       bt.nextID,
		Amount:   amount,
		Category: category,
		Date:     time.Now(),
		Type:     transactionType,
	}
	bt.Transactions = append(bt.Transactions, newTransaction)
	bt.nextID++
}

func (bt BudgetTracker) DisplayTransactions() {
	fmt.Println("ID\tAmount\tCategory\tDate\tType")

	for _, transaction := range bt.Transactions {
		fmt.Printf("%d\t%.2f\t%s\t%s\t%s\n", transaction.ID, transaction.Amount, transaction.Category, transaction.Date.Format("2006-01-02"), transaction.Type)
	}
}

func (bt BudgetTracker) CalculateTotal(tType string) float64 {
	var total float64
	for _, transaction := range bt.Transactions {
		if transaction.Type == tType {
			total += transaction.Amount
		}
	}
	return total
}

// save transactions to csv file
func (bt BudgetTracker) SaveToCSV(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Amount", "Category", "Date", "Type"})

	for _, t := range bt.Transactions {
		record := []string{
			strconv.Itoa(t.ID),
			fmt.Sprintf("%.2f", t.Amount),
			t.Category,
			t.Date.Format("2006-01-02"),
			t.Type,
		}
		err := writer.Write(record)
		if err != nil {
			return err
		}

	}
	fmt.Println("transactions saved to", filename)
	return nil
}

func main() {
	bt := BudgetTracker{}
	for {
		fmt.Println("\n--- Personal Budget Tracker ----")
		fmt.Println("1. Add Transaction")
		fmt.Println("2. Display Transactions")
		fmt.Println("3. Show Total Income")
		fmt.Println("4. Show Total Expenses")
		fmt.Println("5. Save Transactions to CSV")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter amount: ")
			var amount float64
			fmt.Scanln(&amount)

			fmt.Print("Enter category: ")
			var category string
			fmt.Scanln(&category)

			fmt.Print("Enter type (income/expenses): ")
			var tType string
			fmt.Scanln(&tType)

			bt.AddTransaction(amount, category, tType)
			fmt.Println("Transaction added successfully")
		case 2:
			bt.DisplayTransactions()
		case 3:
			fmt.Println("Total Income: ", bt.CalculateTotal("income"))
		case 4:
			fmt.Println("Total Expenses: ", bt.CalculateTotal("expenses"))
		case 5:
			fmt.Print("Enter filename (e.g. transactions.csv): ")
			var filename string
			fmt.Scanln(&filename)
			err := bt.SaveToCSV(filename)
			if err != nil {
				fmt.Println("Error saving transactions:", err)
			}
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")

		}
	}
}
