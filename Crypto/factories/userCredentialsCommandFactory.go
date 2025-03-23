package factories

import (
	"password/commands"
	"password/constants"
	"password/passwords"
)

func CraftUserCredentialsCommand(command string, userName string, password string, verifer *passwords.PasswordVerifier, hasher *passwords.PasswordHasher) commands.Command {
	if command == constants.LOGN_COMMAND {
		return commands.NewLoginCommand(userName, password, verifer)
	} else if command == constants.REGISTER_COMMAND {
		return commands.NewRegisterCommand(userName, password, hasher)
	}
	return nil
}
