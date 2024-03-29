package main

import "strconv"

type currentAccount struct {
	accountNumber int
	clientId      int
	balance       int
}

func (account *currentAccount) GetBalance() int {
	return account.balance
}

func (account *currentAccount) AddBalance(amount int) {
	account.balance += amount
}

func (account *currentAccount) GetAccountDetails() string {
	var accountDetails = "Current Account Number: " + strconv.Itoa(account.accountNumber) +
		"\nCurrent Account Client ID: " + strconv.Itoa(account.clientId) +
		"\nCurrent Account Balance: " + strconv.Itoa(account.GetBalance())

	return accountDetails
}
