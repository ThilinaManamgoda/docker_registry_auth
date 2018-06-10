package main

import (
	"encoding/json"
	"os"
	"github.com/cesanta/docker_auth/auth_server/authz"
	"fmt"
	utils "../utils"
)

const  PushKeyWord  = "push"

func main() {
	text := utils.ReadStdIn()
	// Create the authReqInfo object from the input
	var authReqInfo authz.AuthRequestInfo
	err := json.Unmarshal([]byte(text), &authReqInfo)
	if err != nil {
		os.Exit(utils.ErrorExitCode)
	}

	// Only allowed to "Pull". If "Push" access needed, define the rules via static ACL
	if utils.ArrayContains(authReqInfo.Actions, PushKeyWord) {
		fmt.Println("The user " + authReqInfo.Account + " requesting \"push\" access for the Repo: " + authReqInfo.Name)
		os.Exit(utils.ErrorExitCode)
	}

	repo := authReqInfo.Name
	user := authReqInfo.Account

	isAuthorized := false
	if repo == "hello-world" && user == "admin" {
		isAuthorized = true
	}

	if isAuthorized {
		os.Exit(utils.SuccessExitCode)
	} else {
		os.Exit(utils.ErrorExitCode)
	}
}
