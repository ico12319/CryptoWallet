package users

import "password/passwords"

type UserRepository interface {
	RegisterUser(userName string, password string, hasher *passwords.PasswordHasher) error
	ContainsUser(userName string, password string, verifier *passwords.PasswordVerifier) error
}
