package context

import (
	"context"
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if fetched, err := store.Fetch(r.Context()); err == nil {
			fmt.Fprint(w, fetched)
		}
	}
}

type Store interface {
	Fetch(ctx context.Context) (string, error)
}
