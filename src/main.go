package main

import "fmt"

func main() {

	// Creating a new client based on details provided
	my_client, error := CreateClient("Jorge", "Quintana", 015123172534, "Börnestraße 6, 13086, Berlin")
	if error != nil {
		fmt.Println("Error creating the new user!")
	}
	my_client.PrintClient()

	// Creating Current Account for new client
	my_current_account, error := CreateAccount(my_client)
	if error != nil {
		fmt.Println("Error creating Current Account for the new user!")
	}
	my_current_account.GetAccountDetails()
	my_current_account.AddBalance(100)

	//New Account Details after topping up
	fmt.Println("Topping up the Current Account...")
	fmt.Printf("New Balance: %d", my_current_account.GetBalance())
	fmt.Println()
}
