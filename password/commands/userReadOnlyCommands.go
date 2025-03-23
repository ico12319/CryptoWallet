package commands

import "password/users"

type UserReadOnlyCommands interface {
	GetSummary(user *users.User)
}
