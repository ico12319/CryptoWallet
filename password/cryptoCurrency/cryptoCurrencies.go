package cryptoCurrency

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CryptoCurrencies struct {
	currencies   []CryptoCurrency
	tokensMapped map[string]*CryptoCurrency
}

func NewCryptoCurrencies(resp *http.Response) (*CryptoCurrencies, error) {
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("not ok status code %v\n", resp.Status)
	}
	var cryptoCurr []CryptoCurrency
	err := json.NewDecoder(resp.Body).Decode(&cryptoCurr)
	if err != nil {
		return nil, err
	}
	tempMap := make(map[string]*CryptoCurrency)
	for _, token := range cryptoCurr {
		tempMap[token.AssetId] = &token
	}
	return &CryptoCurrencies{currencies: cryptoCurr, tokensMapped: tempMap}, nil
}

func (c *CryptoCurrencies) ShowListings() {
	for _, cryptoAsset := range c.currencies {
		if cryptoAsset.IsCrypto == 1 {
			cryptoAsset.ShowListing()
		}
	}
}

func (c *CryptoCurrencies) GetMappedTokens() map[string]*CryptoCurrency {
	return c.tokensMapped
}
