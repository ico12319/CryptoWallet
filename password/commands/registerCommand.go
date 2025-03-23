package commands

import (
	"password/passwords"
	"password/users"
)

type RegisterCommand struct {
	userName string
	password string
	hasher   *passwords.PasswordHasher
}

func NewRegisterCommand(userName string, password string, hasher *passwords.PasswordHasher) *RegisterCommand {
	return &RegisterCommand{userName: userName, password: password, hasher: hasher}
}

func (regsiter *RegisterCommand) HandleCommand(dataBase users.UserRepository) error {
	return dataBase.RegisterUser(regsiter.userName, regsiter.password, regsiter.hasher)
}
