package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/Ocelani/swapi-planets/internal/swapi"
	"github.com/Ocelani/swapi-planets/pkg/planet"
)

// main entrypoint.
func main() {
	idFlag := flag.Int("id", 0, "find a planet by ID")
	nameFlag := flag.String("name", "", "find a planet by name")
	flag.Parse()

	planets := swapi.GetPlanets()

	switch {
	case *idFlag != 0:
		id := strconv.Itoa(*idFlag)
		err := planet.FindPlanetID(planets, id)
		fmt.Println(planets, err)

	case *nameFlag != "":
		err := planet.FindPlanetName(planets, *nameFlag)
		fmt.Println(planets, err)

	default:
		fmt.Println(len(planets))
	}
}
