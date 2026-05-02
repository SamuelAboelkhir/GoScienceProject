package commands

import "github.com/SamuelAboelkhir/GoScienceProject/internal/config"

type Command interface {
	Execute(cfg *config.Config, args ...string) error
	Name() string
	Help() string
}

type Commands struct {
	RegisteredCommands map[string]Command
}
