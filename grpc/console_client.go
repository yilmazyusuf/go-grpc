package main

import (
	"fmt"
	"os"

	"./client"
)

func main() {

	if validateUrlArg() == false {
		fmt.Println("Arg Url Required")
		return
	}

	articleUrl := os.Args[1]
	var clientResponse = client.SendRequest(articleUrl)
	fmt.Print(clientResponse)
	fmt.Print("\n\n")

}

func validateUrlArg() bool {
	if len(os.Args) != 2 {
		return false
	}
	return true

}
