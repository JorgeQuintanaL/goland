package main

import "math/rand"

type Account interface {
	GetAccountDetails() string
	GetBalance() int
	AddBalance() int
}

func createAccount(client *personEntity) (*currentAccount, error) {
	return &currentAccount{
		accountNumber: int(rand.Intn(100)),
		clientId:      client.id,
		balance:       0,
	}, nil
}
