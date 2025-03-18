package main

import (
	"fmt"
	"net/http"
	"password/apiKey"
	"password/cryptoCurrency"
	"password/users"
)

func main() {
	req, err := http.NewRequest("GET", apiKey.ASSET_ENDPOINT+"/BTC", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-CoinAPI-Key", apiKey.API_KEY)

	client := &http.Client{}
	resp, err := client.Do(req)

	curr, err := cryptoCurrency.NewCryptoCurrency(resp)
	if err != nil {
		panic(err)
	}
	//curr.ShowListings()

	user := users.NewUser("salamenko", "ivanski04", 100_000)
	user.Buy(1.2, curr)
	user.PrintBalance()
	user.ShowUserAssets()
	fmt.Println()
	user.Sell(1, curr)
	user.PrintBalance()
	user.ShowUserAssets()

}
