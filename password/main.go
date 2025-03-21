package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"password/constants"
	"password/factories"
	"password/passwords"
	"password/users"
	"strings"
)

func main() {
	db, err := sql.Open("sqlite3", "mydb.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	createStatement := `
CREATE TABLE IF NOT EXISTS users (
		username TEXT PRIMARY KEY,
		hashed_password TEXT
	)`

	_, err = db.Exec(createStatement)

	if err != nil {
		panic(err)
	}

	userDb := users.GetInstance(db)
	hasher := passwords.NewPasswordHasher(10)
	verifer := passwords.NewPasswordVerifier()

	reader := bufio.NewReader(os.Stdin)
	userName, _ := reader.ReadString('\n')
	userName = strings.TrimSpace(userName)

	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	command := factories.CraftUserCredentialsCommand(constants.LOGN_COMMAND, userName, password, verifer, hasher)
	fmt.Println(command.HandleCommand(userDb))
}
