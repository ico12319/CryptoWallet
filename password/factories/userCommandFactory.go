package factories

import (
	"bufio"
	"fmt"
	"password/constants"
	"password/priceCache"
	"password/userCommands"
	"strconv"
	"strings"
)

func readInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

func showInfo(endline string) {
	fmt.Printf("Enter %s:\n ", endline)
}

func handleBuySellRequest(reader *bufio.Reader) (string, float64) {
	showInfo("assetId")
	assetId := readInput(reader)
	showInfo("amount")
	amount := readInput(reader)
	parsedAmount, _ := strconv.ParseFloat(amount, 64)
	return assetId, parsedAmount
}

func handleAddFundsRequest(reader *bufio.Reader) float64 {
	showInfo("amount")
	amount := readInput(reader)
	parsedAmount, _ := strconv.ParseFloat(amount, 64)
	return parsedAmount
}

func handleShowWalletOverviewRequest(reader *bufio.Reader) string {
	showInfo("assetId")
	assetId := readInput(reader)
	return assetId
}

func CraftUserCommand(userCommand string, cache *priceCache.PriceCache, reader *bufio.Reader) (userCommands.UserCommands, error) {
	if userCommand == constants.SELL_TOKEN_OPTION {
		assetId, amount := handleBuySellRequest(reader)
		sellCommand, err := userCommands.NewSellCommand(assetId, amount, cache)
		if err != nil {
			return nil, err
		}
		return sellCommand, nil
	} else if userCommand == constants.BUY_TOKEN_OPTION {
		assetId, amount := handleBuySellRequest(reader)
		buyCommand, err := userCommands.NewBuyCommand(assetId, amount, cache)
		if err != nil {
			return nil, err
		}
		return buyCommand, nil
	} else if userCommand == constants.ADD_FUNDS_OPTION {
		amount := handleAddFundsRequest(reader)
		return userCommands.NewAddFundsCommand(amount), nil
	} else if userCommand == constants.SHOW_CURRENT_BALANCE_OPTION {
		return userCommands.NewCurrentBalanceCommand(), nil
	} else if userCommand == constants.SHOW_WALLET_OVERVIEW {
		assetId := handleShowWalletOverviewRequest(reader)
		return userCommands.NewWalletOverview(assetId, cache), nil
	} else if userCommand == constants.SHOW_PORTFOLIO_OPTION {
		return userCommands.NewPortfolioCommand(), nil
	}
	return nil, fmt.Errorf("invalid user command")
}
