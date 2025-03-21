package main

import (
	"bufio"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"password/engine"
)

func main() {
	eng := engine.NewEngine()
	reader := bufio.NewReader(os.Stdin)
	eng.Start(reader)

}
