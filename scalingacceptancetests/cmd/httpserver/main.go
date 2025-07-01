package main

import (
	httpadapters "github.com/gnori-zon/go-tdd/scalingacceptancetests/adapters/httpserver"
	"log"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":8080", httpadapters.NewHandler()); err != nil {
		log.Fatal(err)
	}
}
