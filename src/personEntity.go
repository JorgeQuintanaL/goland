package main

import (
	"strconv"
)

type personEntity struct {
	id      int
	name    string
	phone   int
	address string
}

func (client *personEntity) populateEntity(cl Entity, id int, name string, phone int, address string) error {
	client.id = id
	client.name = name
	client.phone = phone
	client.address = address
	return nil
}

func (client *personEntity) getEntityDetails() string {

	var details = "Client ID: " + strconv.Itoa(client.id) +
		"\nClient Name: " + client.name +
		"\nClient Phone: " + strconv.Itoa(client.phone) +
		"\nClient Address: " + client.address

	return details
}
