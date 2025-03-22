package userBuySellCommands

import (
	"password/apiCaller"
	"password/priceCache"
	"password/users"
)

type UserActionSell struct {
	assetId      string
	amountToSell float64
	cache        *priceCache.PriceCache
	updater      *apiCaller.ApiCaller
}

func NewActionSell(assetId string, amountToSell float64, cache *priceCache.PriceCache, updater *apiCaller.ApiCaller) *UserActionSell {
	return &UserActionSell{assetId: assetId, amountToSell: amountToSell, cache: cache, updater: updater}
}

func (uASell *UserActionSell) HandleActionWithToken(user *users.User) {
	user.Sell(uASell.assetId, uASell.amountToSell, uASell.cache, uASell.updater)
}
