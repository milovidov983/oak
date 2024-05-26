package main

import (
	"fmt"
	"os"
)

type UserInput struct {
	IsResource   bool
	IsValid      bool
	ResourceName string
}
type Config struct {
	IsResource bool
}

func main() {
	firstArg := os.Args[1]

	userInput := createUserInput(os.Args)
	if !userInput.IsValid {
		fmt.Println("Invalid input")
		return
	}

	config := createConfig()
	if userInput.IsResource {
		resource := createResource(userInput.ResourceName, config)
	}

}

func createUserInput(s []string) UserInput {
	return UserInput{
		IsResource: true,
		IsValid:    true,
	}
}

func createConfig() Config {
	panic("unimplemented")
}
