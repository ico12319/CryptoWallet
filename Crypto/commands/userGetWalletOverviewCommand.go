package commands

import (
	"password/priceCache"
	"password/users"
)

type UserWalletOverviewCommand struct{}

func NewUserWalletOverviewCommand() *UserWalletOverviewCommand {
	return &UserWalletOverviewCommand{}
}

func (this *UserWalletOverviewCommand) GetSummaryUsingCache(user *users.User, cacher *priceCache.PriceCache) {
	user.GetWalletOverallSummary(cacher)
}
