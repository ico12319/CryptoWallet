package engine

import (
	"bufio"
	"fmt"
	"password/constants"
	"password/factories"
	"password/helpers"
	"password/passwords"
	"password/priceCache"
	"password/runner"
	"password/users"
)

type Engine struct{}

func NewEngine() *Engine {
	return &Engine{}
}

func (engine *Engine) Start(reader *bufio.Reader) error {
	sqlDatabase, err := helpers.OpenDatabase()
	if err != nil {
		return err
	}
	defer sqlDatabase.Close()

	usersDatabase := users.GetInstance(sqlDatabase)
	cachedPrices := priceCache.GetInstance()

	helpers.ShowWelcomeMessage()
	var command runner.Command

	passwordHasher := passwords.NewPasswordHasher(10) // default value
	passwordVerifier := passwords.NewPasswordVerifier()
	var userName string
	var password string

	for {
		option, err := helpers.SelectOption(reader)
		if err != nil {
			return err
		}
		fmt.Print("Enter username: ")
		userName, err = helpers.SelectOption(reader)
		if err != nil {
			return err
		}
		fmt.Print("Enter password: ")
		password, err = helpers.SelectOption(reader)
		if err != nil {
			return err
		}
		if option == constants.YES_OPTION {
			command = factories.CraftUserCredentialsCommand(constants.LOGN_COMMAND, userName, password, passwordVerifier, passwordHasher)
		} else if option == constants.NO_OPTION {
			command = factories.CraftUserCredentialsCommand(constants.REGISTER_COMMAND, userName, password, passwordVerifier, passwordHasher)
		}
		if command.HandleCommand(usersDatabase) {
			fmt.Printf("Welcome crypto king %s\n", userName)
			break
		} else {
			fmt.Println("Sorry there has been an error trying to process you request! Please try again!")
			continue
		}
	}
	loggedUser := users.NewUser(userName, password, 0)

	for {
		helpers.PrintLoggedUserOptions()
		fmt.Print("Select option: ")
		userOption, err := helpers.SelectOption(reader)
		if err != nil {
			return err
		}
		if userOption == constants.EXIT_OPTION {
			break
		}

		commandObject, err := factories.CraftUserCommand(userOption, cachedPrices, reader)
		commandObject.HandleUserCommand(loggedUser)
	}
	return nil
}
