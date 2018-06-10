package main

import (
	"strings"
	"os"
	"fmt"
	utils "../utils"
)



func main() {
	text := utils.ReadStdIn()
	credentials := strings.Split(text, " ")

	if len(credentials) != 2 {
		fmt.Println("Cannot parse the Input from the Auth service")
		os.Exit(utils.ErrorExitCode)
	}
	uName := credentials[0]
	password := credentials[1]

	isUserAuthenticated := false
	if uName == "admin" && password == "admin" {
		isUserAuthenticated = true
	}


	if isUserAuthenticated {
		os.Exit(utils.SuccessExitCode)
	} else {
		os.Exit(utils.ErrorExitCode)
	}
}