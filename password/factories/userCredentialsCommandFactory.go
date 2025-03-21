package factories

import (
	"password/constants"
	"password/passwords"
	"password/runner"
)

func CraftUserCredentialsCommand(command string, userName string, password string, verifer *passwords.PasswordVerifier, hasher *passwords.PasswordHasher) runner.Command {
	if command == constants.LOGN_COMMAND {
		return runner.NewLoginCommand(userName, password, verifer)
	} else if command == constants.REGISTER_COMMAND {
		return runner.NewRegisterCommand(userName, password, hasher)
	}
	return nil
}
