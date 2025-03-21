package main

import (
	_ "github.com/mattn/go-sqlite3"
	"password/constants"
	"password/factories"
	"password/priceCache"
	"password/users"
)

func main() {
	cache := priceCache.GetInstance()
	user := users.NewUser("ivan", "", 0)
	command1, _ := factories.CraftUserCommand(constants.ADD_FUNDS_OPTION, "0", 100_000, cache)
	command1.HandleUserCommand(user)
	command2, _ := factories.CraftUserCommand(constants.SHOW_CURRENT_BALANCE_OPTION, "", 0, cache)
	command2.HandleUserCommand(user)

	command3, _ := factories.CraftUserCommand(constants.BUY_TOKEN_OPTION, "BTC", 1.1, cache)
	command3.HandleUserCommand(user)

	command4, _ := factories.CraftUserCommand(constants.SHOW_PORTFOLIO_OPTION, "", 0, cache)
	command4.HandleUserCommand(user)

	command5, _ := factories.CraftUserCommand(constants.SELL_TOKEN_OPTION, "BTC", 1, cache)
	command5.HandleUserCommand(user)

	command6, _ := factories.CraftUserCommand(constants.SHOW_PORTFOLIO_OPTION, "", 0, cache)
	command6.HandleUserCommand(user)

}
