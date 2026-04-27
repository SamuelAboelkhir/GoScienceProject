package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func repl(c *config) {
	scanner := bufio.NewScanner(os.Stdin)
	steps := []string{"domain: ", "namespace: ", "identifier: ", "operation (optional): ", "output format: "}

	for {
		fmt.Print("PubChem > ")
		// scanner.Scan()
		// params := cleanInput(scanner.Text())
		// if len(params) == 0 {
		// 	continue
		// }

		query := queryConstructor{}
		for _, step := range steps {
			buildQuery(&query, scanner, step)
			fmt.Println(query)
		}
		elements, err := c.apiClient.API(query.input.domain, query.input.namespace, query.input.identifiers, query.output)
		fmt.Println(elements)
		if err != nil {
			log.Fatal(err)
		}
		for _, element := range elements.PCSubstances {
			fmt.Println(element.Synonyms)
		}

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
