package planet

import "github.com/Ocelani/swapi-planets/gen"

// Repository interface allows us to access the CRUD operations of the database.
type Repository interface {
	Create(*gen.Planet) error
	ReadAll() ([]*gen.Planet, error)
	ReadOne(id string) (*gen.Planet, error)
	Update(*gen.Planet) error
	Delete(id string) error
}

// Service interface methods to access our repository operations.
type Service interface {
	Repository
}
