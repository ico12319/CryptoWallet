package userCommands

import (
	"password/apiCaller"
	"password/priceCache"
	"password/users"
)

type WalletOverviewCommand struct {
	priceUpdater *apiCaller.ApiCaller
}

func NewWalletOverview(assetId string, cache *priceCache.PriceCache) *WalletOverviewCommand {
	return &WalletOverviewCommand{priceUpdater: apiCaller.NewApiCallerForSingleAsset(assetId, cache)}
}

func (wOverview *WalletOverviewCommand) HandleUserCommand(user *users.User) error {
	user.GetWalletOverallSummary(wOverview.priceUpdater)
	return nil
}
