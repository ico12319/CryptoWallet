package factories

import (
	"fmt"
	"password/apiCaller"
	"password/commands"
	"password/constants"
	"password/priceCache"
)

func CraftActionWithTokenCommand(actionCommand string, assetId string, amount float64, cacher *priceCache.PriceCache, updater apiCaller.PriceFetcher) (commands.UserActionWithToken, error) {
	if actionCommand == constants.BUY_TOKEN_OPTION {
		return commands.NewActionBuy(assetId, amount, cacher, updater), nil
	} else if actionCommand == constants.SELL_TOKEN_OPTION {
		return commands.NewActionSell(assetId, amount, cacher, updater), nil
	}
	return nil, fmt.Errorf("invalid action %s\n", actionCommand)
}
