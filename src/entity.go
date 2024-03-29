package main

import "fmt"

type Entity interface {
	populateEntity(client Entity, id int, name string, phone int, address string) error
	getEntityDetails() string
}

func populateEntity(client Entity, id int, name string, phone int, address string) {
	err := client.populateEntity(
		client,
		id,
		name,
		phone,
		address,
	)
	if err != nil {
		fmt.Println("Error creating the Entity!")
	}
}

func getEntityDetails(client Entity) string {
	return client.getEntityDetails()
}
