package commands

import (
	"password/apiCaller"
	"password/priceCache"
	"password/users"
)

type UserActionBuy struct {
	assetId     string
	amountToBuy float64
	cache       *priceCache.PriceCache
	updater     apiCaller.PriceFetcher
}

func NewActionBuy(assetId string, amountToBuy float64, cache *priceCache.PriceCache, updater apiCaller.PriceFetcher) *UserActionBuy {
	return &UserActionBuy{assetId: assetId, amountToBuy: amountToBuy, cache: cache, updater: updater}
}

func (uABuy *UserActionBuy) HandleActionWithToken(user *users.User) error {
	return user.Buy(uABuy.assetId, uABuy.amountToBuy, uABuy.cache, uABuy.updater)
}
