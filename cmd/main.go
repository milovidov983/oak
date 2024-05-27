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

	config, err := createConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	output, err := createOutput(userInput, config)
	if err != nil {
		fmt.Println(err)
		return
	}

	if userInput.IsResource {
		resource, err := findResource(userInput.ResourceName, config)
		if err != nil {
			fmt.Println(err)
			return
		}

		password, err := requestPassword()
		if err != nil {
			fmt.Println(err)
			return
		}
		secret.err := resource.Decode(password)
		if err != nil {
			fmt.Println(err)
			return
		}
		output.Send(secret)

	}

}

func createOutput(userInput UserInput, config Config) (models.Output, error) {
	panic("unimplemented")
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

func createConfig() (Config, error) {
	panic("unimplemented")
}
