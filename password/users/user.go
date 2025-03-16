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
	user.wallet -= amount
}

func (user *User) Buy(assetId string, listings *cryptoCurrency.CryptoCurrencies, amount float64) error {
	desiredCoin := listings.FindByOfferingCode(assetId)
	if desiredCoin == nil {
		return fmt.Errorf("invalid asset id %s\n", assetId)
	}
	if desiredCoin.IsCrypto != 1 {
		return fmt.Errorf("the selected type is not a crypto")
	}
	if desiredCoin.Price > user.wallet || desiredCoin.Price*amount > user.wallet {
		return fmt.Errorf("the user does nto have enough balance to purchase the desired coin")
	}
	user.updateWallet(desiredCoin.Price * amount)
	user.cryptoHoldings[desiredCoin.Name] += amount
	return nil
}

func (user *User) ShowUserAssets() {
	for cryptoToken, quantity := range user.cryptoHoldings {
		fmt.Printf("%s with quantity %0.2f", cryptoToken, quantity)
	}
}

func (user *User) PrintBalance() {
	fmt.Printf("Current balance: %f\n", user.wallet)
}
