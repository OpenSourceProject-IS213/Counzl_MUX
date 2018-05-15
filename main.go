package main

	import (
		"./modules/new_user"
		"fmt"
	)

const (
	P_NEW_USER = "8081"
	P_CHAT = "8082"
)


	func main() {
		fmt.Println("Launching server on:")
		fmt.Println("Port: " + P_NEW_USER)
		new_user.Initialise_User_Listener()
	}
	