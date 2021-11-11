package planet

import "github.com/Ocelani/swapi-planets/gen"

// FindPlanetName finds a planet with the provided name.
// If it doesn't exist, then, it returns 'nil' value.
func FindPlanetName(planets []*gen.Planet, name string) *gen.Planet {
	for _, p := range planets {
		if p.Name == name {
			return p
		}
	}
	return nil
}

// FindPlanetID finds a planet with the provided id.
// If it doesn't exist, then, it returns 'nil' value.
func FindPlanetID(planets []*gen.Planet, id string) *gen.Planet {
	for _, p := range planets {
		if p.Id == id {
			return p
		}
	}
	return nil
}
