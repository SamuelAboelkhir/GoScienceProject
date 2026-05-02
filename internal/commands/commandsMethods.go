package commands

import (
	"errors"

	"github.com/SamuelAboelkhir/GoScienceProject/internal/config"
)

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
