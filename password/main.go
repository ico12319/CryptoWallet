package main

import (
	"net/http"
	"password/apiKey"
	"password/cryptoCurrency"
)

func main() {
	req, err := http.NewRequest("GET", apiKey.ASSET_ENDPOINT+"/ETH", nil)
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
	curr.ShowListing()

	//user := users.NewUser("arbanasenko", "ivanski04", 100_000)
	//cache := priceCache.GetInstance()
	//user.Buy(1, curr, cache)
	//priceUpdate := apiCaller.NewApiCallerForSingleAsset(curr.AssetId, cache)
	//user.GetWalletOverallSummary(priceUpdate)

	//time.Sleep(time.Minute)

	//user.GetWalletOverallSummary(priceUpdate)
	
}
