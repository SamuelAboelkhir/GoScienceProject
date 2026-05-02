package main

import (
	"bufio"
	"os"
	"time"

	"github.com/SamuelAboelkhir/GoScienceProject/internal/pubchemclient"
)

func main() {
	newClient := pubchemclient.NewClient(5*time.Second, 5*time.Minute)
	newScanner := bufio.NewScanner(os.Stdin)
	cfg := config{
		apiClient: newClient,
		scanner:   newScanner,
	}

	c := Commands{make(map[string]func(*config, ...string) error)}

	c.register("buildQuery", queryCommandHandler)

	repl(&cfg, &c)
}
