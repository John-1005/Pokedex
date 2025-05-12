package main


import (
        "fmt"
        "strings"
				"bufio"
				"os"
				"math/rand"
    		"github.com/John-1005/Pokedex/internal/pokeapi"
)

type cliCommand struct{
	name string
	description string
	callback func(*Config) error
}

type Config struct {
	NextURL string
	PreviousURL string
	Args []string
	CaughtPokemon map[string]pokeapi.PokemonDetails
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
			"explore": {
				name: 			 "explore",
				description: "Shows list of pokemon in the location",
				callback:     commandExplore,
			},
			"catch": {
				name: 			 "catch",
				description: "Catch a pokemon",
				callback: 	  commandCatch,
			},
			"information": {
				name:        "information",
				description: "Provides details about caught pokemon",
				callback:     commandInformation,
			},
			"pokedex": {
				name:        "pokedex",
				description: "Lists all caught Pokemon",
				callback:     commandPokedex,
			},
	 }
 }


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := &Config{
			CaughtPokemon: make(map[string]pokeapi.PokemonDetails),
	}
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
		config.Args = words[1:]
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

func commandExplore (config *Config) error {

	if len(config.Args) == 0 {
		fmt.Println("Please specify a location area to explore")
		return nil
	}

	client := pokeapi.NewClient()

	location, err := client.GetLocationArea(config.Args[0])
	if err != nil {
		fmt.Println("invaild location")
		return err
	}
	fmt.Printf("Exploring %s... \n", config.Args[0])
	if len(location.PokemonEncounters) == 0 {
		fmt.Println("expected list of pokemon")
		return nil
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf(" -%s\n", encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch (config *Config) error {
	if len(config.Args) == 0 {
		fmt.Println("Please choose which pokemon you want to catch")
		return nil
	}

	client := pokeapi.NewClient()
	pokemonName := config.Args[0]

	pokemon, err := client.PokemonDetails(pokemonName)
	if err != nil {
		return fmt.Errorf("Failed to get information about %s: %v", pokemonName, err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemonExperience := pokemon.BaseExperience / 2
	randomNumber := rand.Intn(201)

	if randomNumber >= pokemonExperience {
		config.CaughtPokemon[pokemonName] = pokemon
		fmt.Printf("%s was caught!\n", pokemonName)
	}else{
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
}

func commandInformation (config *Config) error {
	pokemonName := config.Args[0]
	pokemonDetails, exists := config.CaughtPokemon[pokemonName]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemonDetails.Name)
	fmt.Printf("Height: %d\n", pokemonDetails.Height)
	fmt.Printf("Weight: %d\n", pokemonDetails.Weight)


	fmt.Println("Stats:")
	for _, stat := range pokemonDetails.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	

	fmt.Println("Types:")
	for _, typeInfo := range pokemonDetails.Types {
		fmt.Printf(" - %s\n", typeInfo.Type.Name)
	}
	return nil
}


func commandPokedex (config *Config) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.CaughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}

