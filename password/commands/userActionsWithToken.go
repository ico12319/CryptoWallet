package commands

import "password/users"

type UserActionWithToken interface {
	HandleActionWithToken(user *users.User) error
}
