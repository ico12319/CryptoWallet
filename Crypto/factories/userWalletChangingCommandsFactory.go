package factories

import (
	"fmt"
	"password/commands"
	"password/constants"
)

func CraftUserWalletChangingCommand(command string) (commands.UserWalletChangingCommand, error) {
	if command == constants.ADD_FUNDS_OPTION {
		return commands.NewUserAddFundsCommand(), nil
	}
	return nil, fmt.Errorf("invalid command %s\n", command)
}
