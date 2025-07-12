package main

import (
	poker "github.com/gnori-zon/go-tdd/app"
	"log"
	"net/http"
)

const dbFileName = "app/game.db.json"

func main() {
	store, close, err := poker.NewFileSystemPlayerStoreFromFile(dbFileName)
	defer close()
	if err != nil {
		log.Fatalf("problem create store %v", err)
	}
	server := poker.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":8080", server))
}
