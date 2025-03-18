package users

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"password/cryptoCurrency"
	"password/passwords"
)

type User struct {
	username       string
	password       string
	wallet         float64
	cryptoHoldings map[string]float64
}

func NewUser(username string, password string, wallet float64) *User {
	tempMap := make(map[string]float64)
	return &User{username: username, password: password, wallet: wallet, cryptoHoldings: tempMap}
}

func (user *User) WriteUserToFile(writer *bufio.Writer) error {
	encrypter, err := passwords.NewPasswordEncrypter()
	if err != nil {
		return err
	}
	_, err = writer.WriteString(user.username + "\n")
	if err != nil {
		return err
	}
	encryptedPassword, err := encrypter.EncryptPassword(user.password)
	if err != nil {
		return err
	}
	encodedPassword := base64.StdEncoding.EncodeToString(encryptedPassword) //converts binary data into a readable text
	_, err = writer.WriteString(encodedPassword + "\n")
	if err != nil {
		return err
	}
	return writer.Flush()
}

func (user *User) DepositMoney(amount float64) {
	user.wallet += amount
}

func (user *User) updateWallet(amount float64) {
	user.wallet += amount
}

func (user *User) Buy(amount float64, token *cryptoCurrency.CryptoCurrency) error {
	if token == nil {
		return fmt.Errorf("invalid asset id %s\n", token.AssetId)
	}
	if token.IsCrypto != 1 {
		return fmt.Errorf("the selected type is not a crypto")
	}
	if token.Price > user.wallet || token.Price*amount > user.wallet {
		return fmt.Errorf("the user does nto have enough balance to purchase the desired coin")
	}
	user.updateWallet(-(token.Price * amount))
	user.cryptoHoldings[token.Name] += amount
	return nil
}

func (user *User) ShowUserAssets() {
	for cryptoToken, quantity := range user.cryptoHoldings {
		fmt.Printf("%s with quantity %0.2f\n", cryptoToken, quantity)
	}
}

func (user *User) PrintBalance() {
	fmt.Printf("Current balance: %f\n", user.wallet)
}

func (user *User) Sell(amount float64, token *cryptoCurrency.CryptoCurrency) error {
	currAmount, contained := user.cryptoHoldings[token.Name]
	if !contained {
		return fmt.Errorf("you don't own a crypto token with such asset id %s\n", token.AssetId)
	}
	updatedAmount := currAmount - amount
	if updatedAmount < 0 {
		return fmt.Errorf("you currently only have %f\n", currAmount)
	}
	if updatedAmount == 0 {
		delete(user.cryptoHoldings, token.AssetId)
	}
	user.cryptoHoldings[token.Name] = updatedAmount
	user.updateWallet(token.Price * amount)
	return nil
}
