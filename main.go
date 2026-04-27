package main

import (
	"time"

	"github.com/SamuelAboelkhir/GoScienceProject/internal/client"
)

func main() {
	newClient := client.NewClient(5*time.Second, 5*time.Minute)
	cfg := config{
		apiClient: newClient,
	}

	repl(&cfg)
}
