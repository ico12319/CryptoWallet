package userCommands

import "password/priceCache"

type CommandParams struct {
	AssetID    string
	Amount     float64
	PriceCache *priceCache.PriceCache
}

func NewCommandParams(cache *priceCache.PriceCache) *CommandParams {
	return &CommandParams{PriceCache: cache}
}

func (cP *CommandParams) SetAssetId(assetId string) {
	cP.AssetID = assetId
}

func (cP *CommandParams) SetAmount(amount float64) {
	cP.Amount = amount
}
