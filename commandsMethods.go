package main

import "errors"

func (c *Commands) run(cfg *config, name string, args ...string) error {
	command, ok := c.registeredCommands[name]
	if !ok {
		return errors.New("please provide a valid command name")
	}
	err := command(cfg, args...)
	if err != nil {
		return err
	}
	return nil
}

func (c *Commands) register(name string, command func(*config, ...string) error) {
	c.registeredCommands[name] = command
}
