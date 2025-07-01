package main

import (
	"github.com/gnori-zon/go-tdd/scalingacceptancetests/adapters/grpcserver"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	grpcserver.RegisterGreeterServer(server, &grpcserver.GreetServer{})

	if err := server.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
