package factories

import (
	"fmt"
	"password/commands"
	"password/constants"
)

func CraftUserCacheNeededCommand(command string) (commands.UserCacheNeededCommands, error) {
	if command == constants.SHOW_WALLET_OVERVIEW {
		return commands.NewUserWalletOverviewCommand(), nil
	}
	return nil, fmt.Errorf("invalid command %s\n", command)
}
