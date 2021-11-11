package planet

import (
	"github.com/Ocelani/swapi-planets/gen"
)

// DefaultService is a service of gen.Planet.
type DefaultService struct{ repo Repository }

// NewPlanet is a constructor of type NewPlanet service.
func NewDefaultService(repo Repository) *DefaultService {
	return &DefaultService{repo}
}

// Create just registers a gen.Planet data on the repository.
func (s *DefaultService) Create(data *gen.Planet) error {
	return s.repo.Create(data)
}

// ReadAll returns the entire data found on this repository.
func (s *DefaultService) ReadAll() ([]*gen.Planet, error) {
	return s.repo.ReadAll()
}

// ReadOne finds and returns the data of a single gen.Planet with provided id.
func (s *DefaultService) ReadOne(id string) (*gen.Planet, error) {
	return s.repo.ReadOne(id)
}

// Update searches the gen.Planet parameter ID, then, updates its data on the repository.
func (s *DefaultService) Update(up *gen.Planet) error {
	save, err := s.ReadOne(up.Id)
	if err != nil {
		return err
	}
	switch {
	case isUpdateField(save.Name, up.Name):
		save.Name = up.Name
		fallthrough

	case isUpdateField(save.Climate, up.Climate):
		save.Climate = up.Climate
		fallthrough

	case isUpdateField(save.Terrain, up.Terrain):
		save.Terrain = up.Terrain
		fallthrough

	case isUpdateField(save.Film, up.Film):
		save.Film = up.Film

	default:
		return nil
	}

	return s.repo.Update(save)
}

// Delete the specific gen.Planet data on the repository with its id as a parameter.
func (s *DefaultService) Delete(id string) error {
	return s.repo.Delete(id)
}

// isUpdateField is used to find if the saved field was updated on the request data.
func isUpdateField(saved, update interface{}) bool {
	switch up := update.(type) {
	case string:
		isEmpty := (up == "")
		isUpdate := (up != saved.(string))
		return !isEmpty && isUpdate

	default:
		return false
	}
}
