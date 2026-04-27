package main

import "github.com/SamuelAboelkhir/GoScienceProject/internal/client"

type config struct {
	apiClient client.Client
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
