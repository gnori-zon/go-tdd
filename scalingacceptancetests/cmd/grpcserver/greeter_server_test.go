package main

import (
	"fmt"
	"github.com/gnori-zon/go-tdd/scalingacceptancetests/adapters"
	"github.com/gnori-zon/go-tdd/scalingacceptancetests/adapters/grpcserver"
	"github.com/gnori-zon/go-tdd/scalingacceptancetests/specification"
	"testing"
)

func TestGreeterServer(t *testing.T) {
	var (
		port       = "50051"
		binToBuild = "grpcserver"
		driver     = &grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	adapters.StartDockerServer(t, port, binToBuild)
	specification.GreetSpecification(t, driver)
	specification.CurseSpecification(t, driver)
}
