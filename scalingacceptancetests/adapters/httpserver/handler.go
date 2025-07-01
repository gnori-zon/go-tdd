package httpserver

import (
	"fmt"
	"github.com/gnori-zon/go-tdd/scalingacceptancetests/domain/intersection"
	"net/http"
)

const (
	greetPath = "/greet"
	cursePath = "/curse"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(greetPath, replyWith(intersection.Greet))
	mux.HandleFunc(cursePath, replyWith(intersection.Curse))
	return mux
}

func replyWith(f func(name string) (interaction string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		fmt.Fprint(w, f(name))
	}
}
