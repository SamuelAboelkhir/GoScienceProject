package main

import (
	"bufio"
	"os"
	"time"

	"github.com/SamuelAboelkhir/GoScienceProject/internal/commands"
	"github.com/SamuelAboelkhir/GoScienceProject/internal/config"
	"github.com/SamuelAboelkhir/GoScienceProject/internal/pubchemclient"
)

func main() {
	newClient := pubchemclient.NewClient(5*time.Second, 5*time.Minute)
	newScanner := bufio.NewScanner(os.Stdin)

	c := commands.Commands{
		RegisteredCommands: make(map[string]commands.Command),
	}

	builder := commands.QueryBuilder{}
	help := commands.Help{
		Commands: &c,
	}
	c.Register(builder.Name(), builder)
	c.Register(help.Name(), help)

	cfg := config.Config{
		APIClient: newClient,
		Scanner:   newScanner,
	}

	repl(&cfg, &c)
}
