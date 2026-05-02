package commands

import (
	"fmt"

	"github.com/SamuelAboelkhir/GoScienceProject/internal/config"
)

type Help struct {
	Commands *Commands
}

func (h Help) Execute(cfg *config.Config, args ...string) error {
	fmt.Println("Welcome to the GoScienceProject CLI!")
	fmt.Println("Available Commands:")
	fmt.Println("-------------------")
	for _, command := range h.Commands.RegisteredCommands {
		fmt.Printf("%s: %s\n\n", command.Name(), command.Help())
	}
	return nil
}

func (h Help) Name() string {
	return "help"
}

func (h Help) Help() string {
	return "Prints this help menu"
}
