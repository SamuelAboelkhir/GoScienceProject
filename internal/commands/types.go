// Package commands: This package is where all the commands are defined. Commands implement a Command interface to allow for more flexibility with their inputs and outputss
package commands

type QueryConstructor struct {
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
