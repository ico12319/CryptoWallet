package users

import (
	"bufio"
	"fmt"
	"password/apiCaller"
	"password/cryptoCurrency"
	"password/passwords"
	"password/priceCache"
)

type User struct {
	username             string
	password             string
	wallet               float64
	cryptoHoldings       map[string]float64 // key: assetId, value: quantity
	cryptoPurchasePrices map[string]float64 // key: assetId, value: average price
}

func NewUser(username string, password string, wallet float64) *User {
	tempMap := make(map[string]float64)
	tempMap2 := make(map[string]float64)
	return &User{username: username, password: password, wallet: wallet, cryptoHoldings: tempMap, cryptoPurchasePrices: tempMap2}
}

func (user *User) WriteUserToFile(writer *bufio.Writer, hasher *passwords.PasswordHasher) (string, error) {
	_, err := writer.WriteString(user.username + "\n")
	if err != nil {
		return "", err
	}
	hashedPassword, err := hasher.HashPassword(user.password)
	if err != nil {
		return "", err
	}
	_, err = writer.WriteString(hashedPassword + "\n")
	if err != nil {
		return "", err
	}
	return hashedPassword, writer.Flush()
}

func (user *User) DepositMoney(amount float64) {
	user.wallet += amount
}

func (user *User) updateWallet(amount float64) {
	user.wallet += amount
}

func (user *User) Buy(amount float64, token *cryptoCurrency.CryptoCurrency, cache *priceCache.PriceCache) error {
	if token == nil {
		return fmt.Errorf("invalid asset id %s\n", token.AssetId)
	}

	if token.IsCrypto != 1 {
		return fmt.Errorf("the selected type is not a crypto")
	}

	if token.Price > user.wallet || token.Price*amount > user.wallet {
		return fmt.Errorf("the user does nto have enough balance to purchase the desired coin")
	}
	_, ok := cache.GetPrice(token.AssetId)
	if !ok {
		cache.SetPrice(token.AssetId, token.Price)
	}
	user.updateWallet(-(token.Price * amount))
	existingQuantity, exists := user.cryptoHoldings[token.AssetId]

	if exists {
		currentAvgPrice := user.cryptoPurchasePrices[token.AssetId]
		totalCost := currentAvgPrice * existingQuantity
		newTotalCost := token.Price * amount
		newQuantity := existingQuantity + amount
		newAvgPrice := (totalCost + newTotalCost) / newQuantity

		user.cryptoHoldings[token.AssetId] = newQuantity
		user.cryptoPurchasePrices[token.AssetId] = newAvgPrice
	} else {
		user.cryptoHoldings[token.AssetId] = amount
		user.cryptoPurchasePrices[token.AssetId] = token.Price
	}
	return nil
}

func (user *User) Sell(amount float64, token *cryptoCurrency.CryptoCurrency, cache *priceCache.PriceCache) error {
	if token == nil {
		return fmt.Errorf("invalid asset id %s\n", token.AssetId)
	}
	currAmount, contained := user.cryptoHoldings[token.Name]
	if !contained {
		return fmt.Errorf("you don't own a crypto token with such asset id %s\n", token.AssetId)
	}
	updatedAmount := currAmount - amount
	if updatedAmount < 0 {
		return fmt.Errorf("you currently only have %f\n", currAmount)
	}
	_, ok := cache.GetPrice(token.AssetId)
	if !ok {
		cache.SetPrice(token.AssetId, token.Price)
	}
	if updatedAmount == 0 {
		delete(user.cryptoHoldings, token.AssetId)
	}
	user.cryptoHoldings[token.AssetId] = updatedAmount
	user.updateWallet(token.Price * amount)
	return nil
}

func (user *User) GetWalletSummary() {
	fmt.Printf("Current balance: %0.2f$\n", user.wallet)
	fmt.Println("Purchased coins:")
	for name, quantity := range user.cryptoHoldings {
		fmt.Printf("%s: %0.2f\n", name, quantity)
	}
}

func (user *User) GetCryptoHoldings() map[string]float64 {
	return user.cryptoHoldings
}

func (user *User) GetWalletOverallSummary(priceUpdater *apiCaller.ApiCaller) {
	overallProfitLoss := 0.0
	for assetId, quantity := range user.cryptoHoldings {
		cachedPrice, ok := priceUpdater.GetCache().GetPrice(assetId)
		if !ok {
			priceUpdater.UpdatePrice()
			cachedPrice, _ = priceUpdater.GetCache().GetPrice(assetId)
		}
		purchasePrice, exist := user.cryptoPurchasePrices[assetId]
		if !exist {
			fmt.Printf("Няма информация за покупната цена на актив с assetId: %s\n", assetId)
			continue
		}
		profitLoss := (cachedPrice - purchasePrice) * quantity
		overallProfitLoss += profitLoss
		fmt.Printf("Asset id %s: quantity %0.2f: purchase price %0.2f: current price %0.2f\n", assetId, quantity, purchasePrice, cachedPrice)
	}

	if overallProfitLoss > 0 {
		fmt.Printf("Congratulations you made some profit! You have earned %0.2f\n", overallProfitLoss)
		return
	}
	fmt.Printf("You are loosing money! You have lost %0.2f\n", overallProfitLoss)
}

func (user *User) GetPassword() string {
	return user.password
}
