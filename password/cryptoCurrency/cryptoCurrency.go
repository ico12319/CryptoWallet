package cryptoCurrency

import (
	"fmt"
)

type CryptoCurrency struct {
	AssetId  string  `json:"asset_id"`
	Name     string  `json:"name"`
	IsCrypto int     `json:"type_is_crypto"`
	Price    float64 `json:"price_usd"`
}

func (c *CryptoCurrency) ShowListing() {
	fmt.Printf("Asset id: %s\n", c.AssetId)
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Price: %f\n", c.Price)
}
