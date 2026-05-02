package main

import "fmt"

func repl(cfg *config, com *Commands) {
	for {
		fmt.Print("PubChem > ")
		cfg.scanner.Scan()

		input := cleanInput(cfg.scanner.Text())
		if len(input) <= 0 {
			fmt.Println("Please provide a command name, or use 'help' for a list of available commands")
			continue
		}
		commandName := input[0]
		args := input[1:]
		err := com.run(cfg, commandName, args...)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
