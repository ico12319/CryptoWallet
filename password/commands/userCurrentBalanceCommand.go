package commands

import (
	"fmt"
	"password/users"
)

type UserCurrentBalanceCommand struct{}

func NewUserCurrentBalanceCommand() *UserCurrentBalanceCommand {
	return &UserCurrentBalanceCommand{}
}

func (this *UserCurrentBalanceCommand) GetSummary(user *users.User) {
	balance := user.GetBalance()
	fmt.Printf("Your current balance is %0.5f$\n", balance)
}
