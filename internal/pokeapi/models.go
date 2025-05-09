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
