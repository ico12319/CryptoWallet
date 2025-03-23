package helpers

import (
	"bufio"
	"database/sql"
	"fmt"
	"password/constants"
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
	fmt.Println("======================================")
	fmt.Println("       WELCOME TO CRYPTO.COM!         ")
	fmt.Println("======================================")
	fmt.Println("Do you have an account? [yes/no]")
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
	fmt.Println("\n======================================")
	fmt.Println("     MAIN MENU - SELECT AN OPTION     ")
	fmt.Println("======================================")
	fmt.Println("1. Add funds")
	fmt.Println("2. Buy crypto token")
	fmt.Println("3. Sell crypto token")
	fmt.Println("4. Show portfolio")
	fmt.Println("5. Show current balance")
	fmt.Println("6. Show wallet overview")
	fmt.Println("7. Show available tokens")
	fmt.Println("8. Exit")
}

func ReadAndParseAmount(reader *bufio.Reader) float64 {
	fmt.Print("Enter amount: ")
	amount, _ := reader.ReadString('\n')
	amount = strings.TrimSpace(amount)
	amountFloat, _ := strconv.ParseFloat(amount, 64)
	return amountFloat
}

func HandleBuySellCommand(reader *bufio.Reader) (string, float64) {
	fmt.Print("Enter assetId: ")
	assetId, _ := reader.ReadString('\n')
	assetId = strings.TrimSpace(assetId)
	parsedAmount := ReadAndParseAmount(reader)
	return assetId, parsedAmount
}

func ValidateOutput(pattern string, pattern2 string) bool {
	return pattern == pattern2
}

func ValidateYesNoCommand(option *string, reader *bufio.Reader) {
	for {
		*option, _ = SelectOption(reader)
		if !ValidateOutput(*option, constants.NO_OPTION) && !ValidateOutput(*option, constants.YES_OPTION) {
			fmt.Println("Invalid option!")
			continue
		}
		break
	}
}
