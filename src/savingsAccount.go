package main

import "strconv"

type savingsAccount struct {
	accountNumber int
	clientId      int
	balance       int
}

func (account *savingsAccount) GetBalance() int {
	return account.balance
}

func (account *savingsAccount) AddBalance(amount int) {
	account.balance += amount
}

func (account *savingsAccount) GetAccountDetails() string {
	var accountDetails = "Current Account Number: " + strconv.Itoa(account.accountNumber) +
		"\nCurrent Account Client ID: " + strconv.Itoa(account.clientId) +
		"\nCurrent Account Balance: " + strconv.Itoa(account.GetBalance())

	return accountDetails
}
