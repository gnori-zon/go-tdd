package grpcserver

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
)

type Driver struct {
	Addr string

	connectionOnce sync.Once
	conn           *grpc.ClientConn
	client         GreeterClient
}

func (d *Driver) Curse(name string) (string, error) {
	client, err := d.getClient()
	if err != nil {
		return "", err
	}
	cursing, err := client.Curse(context.Background(), &CurseRequest{Name: name})
	if err != nil {
		return "", err
	}
	return cursing.Message, nil
}

func (d *Driver) Greet(name string) (string, error) {
	client, err := d.getClient()
	if err != nil {
		return "", err
	}
	greeting, err := client.Greet(context.Background(), &GreetRequest{Name: name})
	if err != nil {
		return "", err
	}
	return greeting.Message, nil
}

func (d *Driver) getClient() (GreeterClient, error) {
	var err error
	d.connectionOnce.Do(func() {
		d.conn, err = grpc.Dial(d.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		d.client = NewGreeterClient(d.conn)
	})
	return d.client, err
}
