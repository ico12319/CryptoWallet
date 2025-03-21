package userCommands

import "password/users"

type UserCommands interface {
	HandleUserCommand(user *users.User) error
}
