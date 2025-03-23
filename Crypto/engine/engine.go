package engine

import (
	"bufio"
	"fmt"
	"password/apiCaller"
	"password/commands"
	"password/constants"
	"password/factories"
	"password/helpers"
	"password/passwords"
	"password/priceCache"
	"password/users"
)

type Engine struct{}

func NewEngine() *Engine {
	return &Engine{}
}

func (engine *Engine) Start(usersDatabase users.UserRepository, cachedPrices *priceCache.PriceCache, reader *bufio.Reader) error {

	helpers.ShowWelcomeMessage()
	var command commands.Command

	passwordHasher := passwords.NewPasswordHasher(10) // default value
	passwordVerifier := passwords.NewPasswordVerifier()
	var userName string
	var password string
	var option string

	for {
		helpers.ValidateYesNoCommand(&option, reader)
		fmt.Print("Enter username: ")
		userName, err := helpers.SelectOption(reader)
		if err != nil {
			fmt.Printf("%s\n", err)
			continue
		}
		fmt.Print("Enter password: ")
		password, err = helpers.SelectOption(reader)
		if err != nil {
			fmt.Println("Invalid credentials!")
			continue
		}
		if option == constants.YES_OPTION {
			command = factories.CraftUserCredentialsCommand(constants.LOGN_COMMAND, userName, password, passwordVerifier, passwordHasher)
		} else if option == constants.NO_OPTION {
			command = factories.CraftUserCredentialsCommand(constants.REGISTER_COMMAND, userName, password, passwordVerifier, passwordHasher)
		}
		err = command.HandleCommand(usersDatabase)
		if err != nil {
			fmt.Printf("%s\n", err)
			continue
		}
		fmt.Printf("Welcome to crypto.com %s\n", userName)
		break
	}
	loggedUser := users.NewUser(userName, password, 0)

	for {
		helpers.PrintLoggedUserOptions()
		fmt.Print("Select option: ")
		userOption, err := helpers.SelectOption(reader)
		if err != nil {
			fmt.Printf("%s\n", err)
			continue
		}
		if userOption == constants.EXIT_OPTION {
			break
		}
		if userOption == constants.BUY_TOKEN_OPTION || userOption == constants.SELL_TOKEN_OPTION {
			assetId, amount := helpers.HandleBuySellCommand(reader)
			updater := apiCaller.NewApiCaller(assetId, cachedPrices)
			action, err := factories.CraftActionWithTokenCommand(userOption, assetId, amount, cachedPrices, updater)
			if err != nil {
				fmt.Printf("%s\n", err)
				continue
			}
			err = action.HandleActionWithToken(loggedUser)
			if err != nil {
				fmt.Printf("%s\n", err)
				continue
			}
		} else if userOption == constants.ADD_FUNDS_OPTION {
			parsedAmount := helpers.ReadAndParseAmount(reader)
			walletChangingCommand, _ := factories.CraftUserWalletChangingCommand(userOption)
			walletChangingCommand.UpdateWallet(loggedUser, parsedAmount)
		} else if userOption == constants.SHOW_PORTFOLIO_OPTION || userOption == constants.SHOW_CURRENT_BALANCE_OPTION {
			readOnlyCommand, _ := factories.CraftUserReadOnlyCommand(userOption)
			readOnlyCommand.GetSummary(loggedUser)
		} else if userOption == constants.SHOW_WALLET_OVERVIEW {
			cacheNeededCommand, _ := factories.CraftUserCacheNeededCommand(userOption)
			cacheNeededCommand.GetSummaryUsingCache(loggedUser, cachedPrices)
		} else if userOption == constants.SHOW_AVAILBLE_TOKENS {
			updater := apiCaller.NewApiCaller("", cachedPrices)
			updater.UpdatePrices()
			listedTokens := updater.GetTokens()
			listedTokens.ShowListings()
		}
	}
	return nil
}
