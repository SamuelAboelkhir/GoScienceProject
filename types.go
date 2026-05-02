package main

import (
	"bufio"

	"github.com/SamuelAboelkhir/GoScienceProject/internal/pubchemclient"
)

type config struct {
	apiClient pubchemclient.PubChemClient
	scanner   *bufio.Scanner
}

type queryConstructor struct {
	input struct {
		domain string
		// TODO: Expand with <structure search>
		namespace   string
		identifiers string
	}
	operation *string
	// NOTE: JSONP can be a callback. Needs investigation
	output string
}

type Commands struct {
	registeredCommands map[string]func(*config, ...string) error
}
