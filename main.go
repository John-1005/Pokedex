package main


import (
        "fmt"
        "strings"
				"bufio"
				"os"
				"time"
    		"github.com/John-1005/Pokedex/internal/pokeapi"
				"github.com/John-1005/Pokedex/internal/pokecache"
)

type cliCommand struct{
	name string
	description string
	callback func(*Config) error
}

type Config struct {
	NextURL string
	PreviousURL string
}

var pokeCache = pokecache.NewCache(5 * time.Minute)
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
				callback:	 	  commandExit,
			},
			"map": {
				name: 			 "map",
				description: "Displays a location",
				callback:			commandMap,
			},
			"mapb": {
				name:				 "mapb",
				description: "Go back to the previou smap location",
				callback:		  commandMapb,
			},
	 }
 }


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := &Config{}
	client =: pokeapi.NewClient(pokeCache)
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
		err := cmd.callback(config)
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

func commandExit (config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}


func commandHelp (config *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range commandRegistry {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap (config *Config) error {
	client := pokeapi.NewClient()
	
	url := ""
	if config != nil {
		url = config.NextURL
	}
	
	rsp, err := client.GetLocationAreas(url)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	
	for _, area := range rsp.Places {
		fmt.Println(area.PlaceName)
	}
	
	if config != nil {

		if rsp.Next != nil {
			config.NextURL = *rsp.Next
		}else {
			config.NextURL = ""
		}

		if rsp.Previous != nil {
			config.PreviousURL = *rsp.Previous
		}else {
			config.PreviousURL = ""
		}
	}
	return nil
}

func commandMapb (config *Config) error {
	
	if config == nil || config.PreviousURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	client := pokeapi.NewClient()

	url := config.PreviousURL

	rsp, err := client.GetLocationAreas(url)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	
	for _, area := range rsp.Places {
		fmt.Println(area.PlaceName)
	}

	if rsp.Next != nil {
		config.NextURL = *rsp.Next
	}else{
		config.NextURL = ""
	}

	if rsp.Previous != nil {
		config.PreviousURL = *rsp.Previous
	}else{
		config.PreviousURL = ""
	}
	return nil
}
