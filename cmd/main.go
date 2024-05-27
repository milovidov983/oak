package main

import (
	"fmt"
	"os"

	"github.com/milovidov983/oak/pkg/models"
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
	userInput := createUserInput(os.Args)
	if !userInput.IsValid {
		fmt.Println("Invalid input")
		return
	}

	config := createConfig()

	if userInput.IsResource {
		resource, err := findResource(userInput.ResourceName, config)
		if err != nil {
			fmt.Println(err)
			return
		}
		// request password

	}

}

func findResource(s string, config Config) (models.Resource, error) {
	panic("unimplemented")
}

func createResource(s string, config Config) {
	panic("unimplemented")
}

func createUserInput(s []string) UserInput {
	// validate user input
	return UserInput{
		IsResource: true,
		IsValid:    true,
	}
}

func createConfig() Config {
	panic("unimplemented")
}
