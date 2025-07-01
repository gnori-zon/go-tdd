package main

import (
	"github.com/gnori-zon/go-tdd/scalingacceptancetests/adapters"
	"github.com/gnori-zon/go-tdd/scalingacceptancetests/adapters/httpserver"
	"github.com/gnori-zon/go-tdd/scalingacceptancetests/specification"
	"net/http"
	"testing"
	"time"
)

func TestGreetingServer(t *testing.T) {
	var (
		port       = "8080"
		binToBuild = "httpserver"
		baseUrl    = "http://localhost:8080"
		driver     = &httpserver.Driver{BaseURL: baseUrl, Client: &http.Client{Timeout: 1 * time.Second}}
	)

	adapters.StartDockerServer(t, port, binToBuild)
	specification.GreetSpecification(t, driver)
	specification.CurseSpecification(t, driver)
}
