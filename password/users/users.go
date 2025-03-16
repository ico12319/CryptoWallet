package users

import (
	"bufio"
	"fmt"
)

type Users struct {
	users map[string]*User
}

func NewUserDatabase() *Users {
	tempUsers := make(map[string]*User)
	return &Users{users: tempUsers}
}

func (u *Users) RegisterNewUser(userName string, password string, writer *bufio.Writer) error {
	newlyRegisteredUser := NewUser(userName, password, 0)
	err := newlyRegisteredUser.WriteUserToFile(writer)
	if err != nil {
		return err
	}
	u.users[userName] = newlyRegisteredUser
	return nil
}

func (u *Users) ShowRegisteredUsers() {
	for _, user := range u.users {
		fmt.Printf("username: %s\n", user.username)
	}
}

func (u *Users) ContainsUser(userName string) *User {
	user, contained := u.users[userName]
	if !contained {
		return nil
	}
	return user
}
