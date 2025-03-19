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
}

func NewApiCallerForSingleAsset(assetId string, cache *priceCache.PriceCache) *ApiCaller {
	formattedUrl := fmt.Sprintf(apiKey.ASSET_ENDPOINT+"/%s", assetId)
	return &ApiCaller{apiKey: apiKey.API_KEY, url: formattedUrl, cache: cache}
}

func NewApiCaller(cache *priceCache.PriceCache) *ApiCaller {
	return &ApiCaller{apiKey: apiKey.API_KEY, url: apiKey.ASSET_ENDPOINT, cache: cache}
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
