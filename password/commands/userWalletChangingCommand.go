package commands

import "password/users"

type UserWalletChangingCommand interface {
	UpdateWallet(user *users.User, amount float64)
}
