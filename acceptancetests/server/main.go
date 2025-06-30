package main

import (
	"context"
	"fmt"
	"github.com/quii/go-graceful-shutdown"
	"log"
	"net/http"
	"time"
)

func main() {
	var (
		ctx        = context.Background()
		httpServer = &http.Server{Addr: ":8080", Handler: http.HandlerFunc(SlowHandler)}
		server     = gracefulshutdown.NewServer(httpServer)
	)

	if err := server.ListenAndServe(ctx); err != nil {
		// this will typically happen if our responses aren't written before the ctx deadline, not much can be done
		log.Fatalf("uh oh, didn't shutdown gracefully, some responses may have been lost %v", err)
	}

	// hopefully, you'll always see this instead
	log.Println("shutdown gracefully! all responses were sent")
}

func SlowHandler(writer http.ResponseWriter, request *http.Request) {
	time.Sleep(2 * time.Second)
	fmt.Fprint(writer, "Hello, world")
}
