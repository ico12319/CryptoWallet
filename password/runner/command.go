package runner

import (
	"password/users"
)

type Command interface {
	HandleCommand(database users.UserRepository) error
}
