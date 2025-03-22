package users

import (
	"fmt"
	"password/apiCaller"
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

func (user *User) DepositMoney(amount float64) {
	user.wallet += amount
	fmt.Printf("You have successfully deposited %0.2f$\n", amount)
}

func (user *User) updateWallet(amount float64) {
	user.wallet += amount
}

func (user *User) Buy(assetId string, amount float64, cache *priceCache.PriceCache, update apiCaller.PriceFetcher) error {
	tokenPrice, isPriceFresh := cache.GetPrice(assetId)
	if !isPriceFresh {
		update.UpdatePrice()
		tokenPrice, _ = cache.GetPrice(assetId)
	}

	if tokenPrice > user.wallet || tokenPrice*amount > user.wallet {
		return fmt.Errorf("not enough balance to purchase the desired coin\n")
	}

	user.updateWallet(-(tokenPrice * amount))
	existingQuantity, exists := user.cryptoHoldings[assetId]

	if exists {
		currentAvgPrice := user.cryptoPurchasePrices[assetId]
		totalCost := currentAvgPrice * existingQuantity
		newTotalCost := tokenPrice * amount
		newQuantity := existingQuantity + amount
		newAvgPrice := (totalCost + newTotalCost) / newQuantity

		user.cryptoHoldings[assetId] = newQuantity
		user.cryptoPurchasePrices[assetId] = newAvgPrice
	} else {
		user.cryptoHoldings[assetId] = amount
		user.cryptoPurchasePrices[assetId] = tokenPrice
	}
	return nil
}

func (user *User) Sell(assetId string, amount float64, cache *priceCache.PriceCache, updater apiCaller.PriceFetcher) error {
	currAmount, contained := user.cryptoHoldings[assetId]
	if !contained {
		return fmt.Errorf("you don't own a crypto token with such asset id\n")
	}

	updatedAmount := currAmount - amount

	if updatedAmount < 0 {
		return fmt.Errorf("you currently only have %f\n", currAmount)
	}

	tokenPrice, isPriceFresh := cache.GetPrice(assetId)
	if !isPriceFresh {
		updater.UpdatePrice()
		tokenPrice, _ = cache.GetPrice(assetId)
	}
	if updatedAmount == 0 {
		delete(user.cryptoHoldings, assetId)
	} else {
		user.cryptoHoldings[assetId] = updatedAmount
	}
	user.updateWallet(tokenPrice * amount)
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

func (user *User) GetWalletOverallSummary(cacher *priceCache.PriceCache) {
	overallProfitLoss := 0.0
	for assetId, quantity := range user.cryptoHoldings {
		cachedPrice, ok := cacher.GetPrice(assetId)
		if !ok {
			priceUpdater := apiCaller.NewApiCaller(assetId, cacher)
			priceUpdater.UpdatePrice()
			cachedPrice, _ = cacher.GetPrice(assetId)
		}
		purchasePrice, exist := user.cryptoPurchasePrices[assetId]
		if !exist {
			fmt.Printf("There is not information about the price of a token with assetId: %s\n", assetId)
			continue
		}
		profitLoss := (cachedPrice - purchasePrice) * quantity
		overallProfitLoss += profitLoss
		fmt.Printf("Asset id %s: quantity %0.2f: purchase price %0.2f: current price %0.2f\n", assetId, quantity, purchasePrice, cachedPrice)
	}

	if overallProfitLoss > 0 {
		fmt.Printf("Congratulations you made some profit! You have earned %0.2f\n", overallProfitLoss)
		return
	} else if overallProfitLoss < 0 {
		fmt.Printf("You are loosing money! You have lost %0.2f\n", overallProfitLoss)
	}
}

func (user *User) GetBalance() float64 {
	return user.wallet
}
