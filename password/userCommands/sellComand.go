package userCommands

import (
	"fmt"
	"net/http"
	"password/apiKey"
	"password/cryptoCurrency"
	"password/priceCache"
	"password/users"
)

type SellCommand struct {
	amountToBuy float64
	cache       *priceCache.PriceCache
	token       *cryptoCurrency.CryptoCurrency
}

func NewSellCommand(assetId string, amountToBuy float64, cache *priceCache.PriceCache) (*SellCommand, error) {
	formattedUrl := fmt.Sprintf(apiKey.ASSET_ENDPOINT+"/%s", assetId)
	req, err := http.NewRequest("GET", formattedUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-CoinAPI-Key", apiKey.API_KEY)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("not ok http status code %d\n", resp.StatusCode)
	}

	constructedToken, err := cryptoCurrency.NewCryptoCurrency(resp)
	if err != nil {
		return nil, err
	}
	return &SellCommand{amountToBuy: amountToBuy, cache: cache, token: constructedToken}, nil
}

func (sellCommand *SellCommand) HandleUserCommand(user *users.User) error {
	return user.Sell(sellCommand.amountToBuy, sellCommand.token, sellCommand.cache)
}
