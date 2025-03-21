package userCommands

import "password/users"

type CurrentBalanceCommand struct{}

func NewCurrentBalanceCommand() *CurrentBalanceCommand {
	return &CurrentBalanceCommand{}
}

func (cBalance *CurrentBalanceCommand) HandleUserCommand(user *users.User) error {
	user.GetBalance()
	return nil
}
