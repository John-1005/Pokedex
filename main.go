package main


import (
        "fmt"
        "strings"
				"bufio"
				"os"
)

type cliCommand struct{
	name string
	description string
	callback func() error
}

var commandRegistry map[string]cliCommand 

 func init() {
	 commandRegistry = map[string]cliCommand {
		 "help": {
			 	name: 			 "help",
				description: "Displays a help message", 
				callback: 		commandHelp,
		 } ,
		 "exit": {
			 	name: 			 "exit",
				description: "Exit the Pokedex",
				callback:	 	 commandExit,
			},
	 }
 }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Pokedex >")
		scanner.Scan()
		answer := scanner.Text()
		words := cleanInput(answer)

		if len(words) == 0 {
			fmt.Println("Please enter a command.")
			continue
		}
		cmd, ok := commandRegistry[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := cmd.callback()
		if err != nil {
			fmt.Println(err)
		}
		
	}
}


func cleanInput(text string) []string {
  s := strings.TrimSpace(text)
  lowerChars := strings.ToLower(s)
  words := strings.Fields(lowerChars)
  return words
}

func commandExit () error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}


func commandHelp () error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range commandRegistry {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
