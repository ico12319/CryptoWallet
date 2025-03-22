package apiCaller

import (
	"fmt"
	"net/http"
	"password/apiKey"
	"password/cryptoCurrency"
	"password/priceCache"
)

type ApiCaller struct {
	apiKey string
	url    string
	cache  *priceCache.PriceCache
	tokens *cryptoCurrency.CryptoCurrencies
}

func formatUrl(assetId string) string {
	if len(assetId) != 0 {
		url := fmt.Sprintf(apiKey.ASSET_ENDPOINT+"/%s", assetId)
		return url
	}
	return apiKey.ASSET_ENDPOINT
}

func NewApiCaller(assetId string, cache *priceCache.PriceCache) *ApiCaller {
	url := formatUrl(assetId)
	return &ApiCaller{apiKey: apiKey.API_KEY, url: url, cache: cache, tokens: nil}
}

func (aCaller *ApiCaller) UpdatePrice() error {
	req, err := http.NewRequest("GET", aCaller.url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-CoinAPI-Key", aCaller.apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("not ok status code %d\n", resp.StatusCode)
	}

	defer resp.Body.Close()

	if err != nil {
		return err
	}
	token, err := cryptoCurrency.NewCryptoCurrency(resp)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	aCaller.GetCache().SetPrice(token.AssetId, token.Price)
	return nil
}

func (aCaller *ApiCaller) GetCache() *priceCache.PriceCache {
	return aCaller.cache
}

func (aCaller *ApiCaller) UpdatePrices() error {
	req, err := http.NewRequest("GET", aCaller.url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-CoinAPI-Key", aCaller.apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("not okay status code %d\n", resp.StatusCode)
	}
	defer resp.Body.Close()

	if err != nil {
		return err
	}
	aCaller.tokens, err = cryptoCurrency.NewCryptoCurrencies(resp)
	if err != nil {
		return err
	}
	return nil
}

func (aCaller *ApiCaller) GetTokens() *cryptoCurrency.CryptoCurrencies {
	return aCaller.tokens
}
