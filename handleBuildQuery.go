package main

import (
	"bufio"
	"fmt"
)

func buildQuery(query *queryConstructor, scanner *bufio.Scanner, step string) {
	fmt.Printf("Please provide a %s", step)
	scanner.Scan()
	param := scanner.Text()
	switch step {
	case "domain: ":
		query.input.domain = param
	case "namespace: ":
		query.input.namespace = param
	case "identifier: ":
		query.input.identifiers = param
	case "operation (optional): ":
		if len(param) > 0 {
			query.operation = &param
		} else {
			query.operation = nil
		}
	case "output format: ":
		query.output = param
	}
}
