package commands

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

func (logger *LoginCommand) HandleCommand(userDatabase users.UserRepository) error {
	return userDatabase.ContainsUser(logger.userName, logger.password, logger.verifier)
}
