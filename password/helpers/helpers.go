package helpers

import (
	"bufio"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

func OpenDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "mydb.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ShowWelcomeMessage() {
	fmt.Println("Welcome to Crypto.com!")
	fmt.Println("Do you have an account?")
}

func SelectOption(reader *bufio.Reader) (string, error) {
	option, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	option = strings.TrimSpace(option)
	return option, nil
}

func PrintLoggedUserOptions() {
	fmt.Println("1. Add funds")
	fmt.Println("2. Buy crypto token")
	fmt.Println("3. Sell crypto token")
	fmt.Println("4. Show portfolio")
	fmt.Println("5. Show current balance")
	fmt.Println("6. Show wallet overview")
	fmt.Println("7. Exit")
}

func ReadAndParseAmount(reader *bufio.Reader) float64 {
	fmt.Print("Enter amount: ")
	amount, _ := reader.ReadString('\n')
	amount = strings.TrimSpace(amount)
	amountFloat, _ := strconv.ParseFloat(amount, 64)
	return amountFloat
}
