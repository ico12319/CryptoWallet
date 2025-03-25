package cryptoCurrency

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpCryptoToken struct {
	AssetId  string  `json:"asset_id"`
	Name     string  `json:"name"`
	IsCrypto int     `json:"type_is_crypto"`
	Price    float64 `json:"price_usd"`
}

func NewHttpCryptoCurrency(resp *http.Response) (*HttpCryptoToken, error) {
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("not ok http status code %v\n", resp.StatusCode)
	}
	var token []HttpCryptoToken
	err := json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return nil, err
	}
	return &token[0], nil
}

func (c *HttpCryptoToken) ShowListing() {
	fmt.Printf("Asset id: %s\n", c.AssetId)
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Price: %0.15f\n", c.Price)
}
