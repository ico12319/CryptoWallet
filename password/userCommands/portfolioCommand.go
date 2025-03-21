package userCommands

import "password/users"

type PortfolioCommand struct{}

func NewPortfolioCommand() *PortfolioCommand {
	return &PortfolioCommand{}
}

func (portfolio *PortfolioCommand) HandleUserCommand(user *users.User) error {
	user.GetWalletSummary()
	return nil
}
