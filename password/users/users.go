package users

import (
	"database/sql"
	"password/passwords"
	"sync"
)

// Users Database is implemented using a singleton design pattern
type Users struct {
	dataBase *sql.DB
}

var instance *Users
var once sync.Once

func GetInstance(dataBase *sql.DB) *Users {
	once.Do(func() {
		instance = &Users{dataBase: dataBase}
	})
	return instance
}

func (u *Users) RegisterNewUser(userName string, password string, hasher *passwords.PasswordHasher) error {
	hashedPassword, err := hasher.HashPassword(password)
	if err != nil {
		return err
	}
	insertStatement := `INSERT INTO users(username, hashed_password) VALUES(?, ?)`
	_, err = u.dataBase.Exec(insertStatement, userName, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}

func (u *Users) ContainsUser(userName string, password string, verifier *passwords.PasswordVerifier) bool {
	var hashedPassword string
	selectStatement := `SELECT hashed_password FROM users WHERE username = ? LIMIT 1`
	err := u.dataBase.QueryRow(selectStatement, userName).Scan(&hashedPassword)
	if err != nil {
		return false
	}

	arePasswordMatching := verifier.VerifyPassword(hashedPassword, password)
	return arePasswordMatching
}
