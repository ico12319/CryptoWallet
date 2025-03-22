package main

import (
	"bufio"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"password/engine"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	eng := engine.NewEngine()
	eng.Start(reader)

}
