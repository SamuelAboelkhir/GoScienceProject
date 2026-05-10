// Package config: This internal package defines the config of the app
package config

import (
	"bufio"

	"github.com/SamuelAboelkhir/CompGoR/internal/pubchemclient"
)

type Config struct {
	APIClient *pubchemclient.PubChemClient
	Scanner   *bufio.Scanner
}
