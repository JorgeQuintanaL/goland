package main

import (
	"fmt"
	"math/rand"
)

type Client interface {
	PrintClient()
}

type NewClient struct {
	id         int
	first_name string
	last_name  string
	phone      int
	address    string
}

func CreateClient(first_name, last_name string, phone int, address string) (*NewClient, error) {
	return &NewClient{
		id:         int(rand.Intn(100)),
		first_name: first_name,
		last_name:  last_name,
		phone:      phone,
		address:    address,
	}, nil
}

func (client *NewClient) PrintClient() {

	fmt.Printf("Client ID: %d"+
		"\nClient First Name: %s"+
		"\nClient Last Name: %s"+
		"\nClient Phone: %d"+
		"\nClient Address: %s",
		client.id,
		client.first_name,
		client.last_name,
		client.phone,
		client.address,
	)
	fmt.Println()
}
