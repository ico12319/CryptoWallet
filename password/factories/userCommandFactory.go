package factories

import (
	"fmt"
	"password/constants"
	"password/priceCache"
	"password/userCommands"
)

func CraftUserCommand(userCommand string, assetId string, amount float64, cache *priceCache.PriceCache) (userCommands.UserCommands, error) {
	if userCommand == constants.SELL_TOKEN_OPTION {
		sellCommand, err := userCommands.NewSellCommand(assetId, amount, cache)
		if err != nil {
			return nil, err
		}
		return sellCommand, nil
	} else if userCommand == constants.BUY_TOKEN_OPTION {
		buyCommand, err := userCommands.NewBuyCommand(assetId, amount, cache)
		if err != nil {
			return nil, err
		}
		return buyCommand, nil
	} else if userCommand == constants.ADD_FUNDS_OPTION {
		return userCommands.NewAddFundsCommand(amount), nil
	} else if userCommand == constants.SHOW_CURRENT_BALANCE_OPTION {
		return userCommands.NewCurrentBalanceCommand(), nil
	} else if userCommand == constants.SHOW_WALLET_OVERVIEW {
		return userCommands.NewWalletOverview(assetId, cache), nil
	} else if userCommand == constants.SHOW_PORTFOLIO_OPTION {
		return userCommands.NewPortfolioCommand(), nil
	}
	return nil, fmt.Errorf("invalid user command")
}
