package factories

import (
	"fmt"
	"password/commands"
	"password/constants"
)

func CraftUserReadOnlyCommand(command string) (commands.UserReadOnlyCommands, error) {
	if command == constants.SHOW_PORTFOLIO_OPTION {
		return commands.NewUserGetWalletSummaryCommand(), nil
	} else if command == constants.SHOW_CURRENT_BALANCE_OPTION {
		return commands.NewUserCurrentBalanceCommand(), nil
	}
	return nil, fmt.Errorf("invalid command %s\n", command)
}
