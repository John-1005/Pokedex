package pokeapi





type LocationAreaResponse struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Places []LocationArea `json:"results"`
}


type LocationArea struct {
	PlaceName string `json:"name"`
	DetailsURL string `json:"url"`
}

type LocationAreaPokemon struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type PokemonDetails struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	BaseExperience int      `json:"base_experience"`
  Height         int      `json:"height"`
	Weight         int      `json:"weight"`
	Stats          []Stat   `json:"stats"`
	Types          []Type   `json:"types"`
}


type StatDetail struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Stat struct {
	BaseStat int        `json:"base_stat"`
	Effort   int        `json:"effort"`
	Stat     StatDetail `json:"stat"`
}

type TypeDetail struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Type struct {
	Slot int        `json:"slot"`
	Type TypeDetail `json:"type"`
}

