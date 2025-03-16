package main

import (
	"net/http"
	"password/apiKey"
	"password/cryptoCurrency"
	"password/users"
)

func main() {
	req, err := http.NewRequest("GET", apiKey.ASSET_ENDPOINT, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-CoinAPI-Key", apiKey.API_KEY)

	client := &http.Client{}
	resp, err := client.Do(req)

	curr, err := cryptoCurrency.NewCryptoCurrencies(resp)
	if err != nil {
		panic(err)
	}
	//curr.ShowListings()

	user := users.NewUser("chica", "salam", 100)
	err = user.Buy("DOGE", curr, 100)
	if err != nil {
		panic(err)
	}

	user.ShowUserAssets()


}
