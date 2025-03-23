package commands

import "password/users"

type UserGetWalletSummaryCommand struct{}

func NewUserGetWalletSummaryCommand() *UserGetWalletSummaryCommand {
	return &UserGetWalletSummaryCommand{}
}

func (this *UserGetWalletSummaryCommand) GetSummary(user *users.User) {
	user.GetWalletSummary()
}
