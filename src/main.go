package main

import "fmt"

func readData(id *int, name *string, phone *int, address *string) {

	fmt.Println("What is the client's ID:")
	_, err := fmt.Scanln(id)
	if err != nil {
		fmt.Println("Error reading the input!")
	}
	fmt.Println("What is the client's name:")
	_, err = fmt.Scanln(name)
	fmt.Println("What is the client's phone:")
	_, err = fmt.Scanln(phone)
	fmt.Println("What is the client's address:")
	_, err = fmt.Scanln(address)
}

func main() {

	var id, phone int
	var name, address string
	var myNewEntity personEntity

	readData(&id, &name, &phone, &address)
	populateEntity(&myNewEntity, id, name, phone, address)

	//Printing the new user details
	fmt.Println("The client was successfully created and the details are the following:")
	clientDetails := getEntityDetails(&myNewEntity)
	fmt.Println(clientDetails)
}
