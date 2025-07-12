package main

import (
	"fmt"
	poker "github.com/gnori-zon/go-tdd/app"
	"github.com/gnori-zon/go-tdd/app/cli"
	"log"
	"os"
)

const dbFileName = "app/game.db.json"

func main() {
	fmt.Println("Let's play poker!")
	fmt.Println("Type '{Name} wins' to record a win")
	store, close, err := poker.NewFileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("problem create store %v", err)
	}
	defer close()
	game := cli.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
