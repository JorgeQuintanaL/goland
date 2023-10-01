package main

import (
	"fmt"
	"math/rand"
)

type Account interface {
	GetAccountDetails()
	GetBalance()
	AddBalance()
}

type CurrentAccount struct {
	number    int
	client_id int
	balance   int
}

func CreateAccount(client *NewClient) (*CurrentAccount, error) {
	return &CurrentAccount{
		number:    rand.Intn(1000000),
		client_id: client.id,
	}, nil
}

func (account *CurrentAccount) GetBalance() int {
	return account.balance
}

func (account *CurrentAccount) GetAccountDetails() {
	fmt.Printf(
		"Current Account Number: %d"+
			"\nCurrent Account Client ID: %d"+
			"\nCurrent Account Balance: %d",
		account.number,
		account.client_id,
		account.GetBalance(),
	)
	fmt.Println()
}

func (account *CurrentAccount) AddBalance(amount int) {
	account.balance += amount
}
