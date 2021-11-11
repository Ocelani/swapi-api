package planet

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Ocelani/swapi-planets/gen"
)

// CacheRepository is a repository of gen.Planet
// that uses an array to cache runtime data.
type CacheRepository struct{ Data []*gen.Planet }

// NewCacheRepository is a constructor of type CacheRepository repository.
func NewCacheRepository(data []*gen.Planet) *CacheRepository {
	return &CacheRepository{data}
}

// Create just registers a gen.Planet data on the repository.
func (r *CacheRepository) Create(data *gen.Planet) error {
	p := FindPlanetName(r.Data, data.Id)
	if p != nil {
		return fmt.Errorf("planet already exists")
	}
	id := time.Now().UnixMilli()
	data.Id = strconv.Itoa(int(id))
	r.Data = append(r.Data, data)

	return nil
}

// ReadAll returns the entire data found on this repository.
func (r *CacheRepository) ReadAll() ([]*gen.Planet, error) {
	return r.Data, nil
}

// ReadOne finds and returns the data of a single gen.Planet with provided id.
func (r *CacheRepository) ReadOne(id string) (*gen.Planet, error) {
	if planet := FindPlanetID(r.Data, id); planet != nil {
		return planet, nil
	}
	return nil, fmt.Errorf("planet not found with id: %s", id)
}

// Update searches the gen.Planet parameter Id, then, updates its data on the repository.
func (r *CacheRepository) Update(data *gen.Planet) error {
	for i, p := range r.Data {
		if p.Id == data.Id {
			r.Data[i] = data
			return nil
		}
	}
	return fmt.Errorf("planet not found with id: %s", data.Id)
}

// Delete the specific gen.Planet data on the repository with its id as a parameter.
func (r *CacheRepository) Delete(id string) error {
	for i, saved := range r.Data {
		if saved.Id != id {
			continue
		}
		d1 := r.Data[:i]

		if last := len(r.Data) - 1; last >= i {
			d2 := r.Data[i+1:]
			d1 = append(d1, d2...)
		}
		r.Data = d1

		return nil
	}

	return fmt.Errorf("planet not found with id: %s", id)
}
