package runner

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

func (regsiter *RegisterCommand) HandleCommand(dataBase *users.Users) bool {
	if dataBase.RegisterNewUser(regsiter.userName, regsiter.password, regsiter.hasher) != nil {
		return false
	}
	return true
}
