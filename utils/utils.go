package utils

import (
	"bufio"
	"os"
	"strings"
)

const SuccessExitCode  =  0
const ErrorExitCode  = 1

// Read Standard input stream
func ReadStdIn() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1) // remove "\n" from the string.
	return text
}

// Check whether the array contains the given key
// This method is only efficient for arrays with smaller number of elements
func ArrayContains(array []string, key string) bool {
	for _, tmp := range array {
		if tmp == key {
			return true
		}
	}
	return false
}
