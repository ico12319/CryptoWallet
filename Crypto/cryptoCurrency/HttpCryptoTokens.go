package cryptoCurrency

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpCryptoTokens struct {
	tokens []HttpCryptoToken
}

func NewHttpCryptoCurrencies(resp *http.Response) (*HttpCryptoTokens, error) {
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("not ok status code %v\n", resp.Status)
	}
	var cryptoCurr []HttpCryptoToken
	err := json.NewDecoder(resp.Body).Decode(&cryptoCurr)
	if err != nil {
		return nil, err
	}
	return &HttpCryptoTokens{tokens: cryptoCurr}, nil
}

func (c *HttpCryptoTokens) ShowListings() {
	for _, cryptoAsset := range c.tokens {
		if cryptoAsset.IsCrypto == 1 {
			cryptoAsset.ShowListing()
		}
	}
}
