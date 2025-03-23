package commands

import (
	"fmt"
	"password/users"
)

type UserAddFundsCommand struct{}

func NewUserAddFundsCommand() *UserAddFundsCommand {
	return &UserAddFundsCommand{}
}

func (this *UserAddFundsCommand) UpdateWallet(user *users.User, amount float64) {
	user.DepositMoney(amount)
	fmt.Printf("You have successfully deposited %0.5f$\n", amount)
}
