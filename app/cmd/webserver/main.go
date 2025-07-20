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
	alerter := poker.BlindAlerterFunc(poker.Alerter)
	server, err := poker.NewPlayerServer(store, poker.NewGame(store, alerter))
	if err != nil {
		log.Fatalf("problem create server %v", err)
	}
	log.Fatal(http.ListenAndServe(":8080", server))
}
