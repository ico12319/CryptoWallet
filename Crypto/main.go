package main

import (
	"bufio"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"password/engine"
	"password/helpers"
	"password/priceCache"
	"password/users"
)

func main() {
	dB, err := helpers.OpenDatabase()
	if err != nil {
		panic(err)
	}

	dataBase := users.GetInstance(dB)
	cache := priceCache.GetInstance()
	reader := bufio.NewReader(os.Stdin)

	eng := engine.NewEngine()
	err = eng.Start(dataBase, cache, reader)
	if err != nil {
		panic(err)
	}
}
