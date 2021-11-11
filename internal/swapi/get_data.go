package swapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Ocelani/swapi-planets/gen"
)

// _swapiURL is the base url of the third party API.
const (
	_swapiURL     = "https://swapi.dev/api"
	_planetsRoute = "/planets"
)

// GetPlanets receives the planet.Planet from SWAPI.
func GetPlanets() []*gen.Planet {
	resp, err := http.Get(_swapiURL + _planetsRoute)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	var planets struct {
		Count    int           `json:"count"`
		Next     string        `json:"next"`
		Previous string        `json:"previous"`
		Results  []*gen.Planet `json:"results"`
	}
	if err = json.Unmarshal(body, &planets); err != nil {
		panic(err)
	}

	return planets.Results
}
