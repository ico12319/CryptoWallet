package runner

import (
	"password/passwords"
	"password/users"
)

type LoginCommand struct {
	userName string
	password string
	verifier *passwords.PasswordVerifier
}

func NewLoginCommand(userName string, password string, verifier *passwords.PasswordVerifier) *LoginCommand {
	return &LoginCommand{userName: userName, password: password, verifier: verifier}
}

func (logger *LoginCommand) HandleCommand(userDatabase *users.Users) bool {
	if userDatabase.ContainsUser(logger.userName, logger.password, logger.verifier) {
		return true
	}
	return false
}
