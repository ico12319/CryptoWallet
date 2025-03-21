package userCommands

import "password/users"

type AddFundsCommand struct {
	amountToAdd float64
}

func NewAddFundsCommand(amountToAdd float64) *AddFundsCommand {
	return &AddFundsCommand{amountToAdd: amountToAdd}
}

func (addFunds *AddFundsCommand) HandleUserCommand(user *users.User) error {
	user.DepositMoney(addFunds.amountToAdd)
	return nil
}
