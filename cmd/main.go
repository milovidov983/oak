package main

import (
	"fmt"
	"os"
)

type Command struct {
	Name        string
	Description string
	Action      func([]string)
}
type Result struct {
	Success bool
	Message string
}

func main() {
	command, err := parseCommand(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(command)
	result, err := executeCommand(command)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func executeCommand(command Command) {
	panic("unimplemented")
}

func parseCommand(s []string) Command {
	return Command{Name: "test", Description: "test", Action: test}
}
