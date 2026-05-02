package main

import "fmt"

func queryCommandHandler(cfg *config, args ...string) error {
	fmt.Println("arrived at builder")
	fmt.Println(args)
	steps := []string{"domain: ", "namespace: ", "identifier: ", "operation (optional): ", "output format: "}
	query := queryConstructor{}
	for _, step := range steps {
		buildQuery(cfg, &query, step)
		fmt.Println(query)
	}

	elements, err := cfg.apiClient.GetCompounds(query.input.domain, query.input.namespace, query.input.identifiers, query.output)
	// fmt.Println(elements)
	if err != nil {
		return err
	}
	for _, element := range elements.PCCompounds {
		fmt.Println("Element ID: ", element.ID)
		fmt.Println("Props: ", element.Props)
		for _, atom := range element.Atoms.Element {
			fmt.Println("Atom: ", atom)
		}
	}
	return nil
}
