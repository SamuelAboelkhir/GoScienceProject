// Package commands: This package is where all the commands are defined. Commands implement a Command interface to allow for more flexibility with their inputs and outputss
package commands

import (
	"errors"

	"github.com/SamuelAboelkhir/GoScienceProject/internal/config"
)

type Command interface {
	Execute(cfg *config.Config, args ...string) error
	Name() string
	Help() string
}

type Commands struct {
	RegisteredCommands map[string]Command
}

func (c *Commands) Run(cfg *config.Config, name string, args ...string) error {
	command, ok := c.RegisteredCommands[name]
	if !ok {
		return errors.New("please provide a valid command name")
	}
	err := command.Execute(cfg, args...)
	if err != nil {
		return err
	}
	return nil
}

func (c *Commands) Register(name string, cmd Command) {
	c.RegisteredCommands[name] = cmd
}
