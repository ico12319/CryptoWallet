package factories

import (
	"fmt"
	"password/apiCaller"
	"password/constants"
	"password/priceCache"
	"password/userBuySellCommands"
)

func CraftActionWithTokenCommand(actionCommand string, assetId string, amount float64, cacher *priceCache.PriceCache, updater *apiCaller.ApiCaller) (userBuySellCommands.UserActionWithToken, error) {
	if actionCommand == constants.BUY_TOKEN_OPTION {
		return userBuySellCommands.NewActionBuy(assetId, amount, cacher, updater), nil
	} else if actionCommand == constants.SELL_TOKEN_OPTION {
		return userBuySellCommands.NewActionSell(assetId, amount, cacher, updater), nil
	}
	return nil, fmt.Errorf("invalid action %s\n", actionCommand)
}
