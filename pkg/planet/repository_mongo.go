package planet

import (
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Ocelani/swapi-planets/gen"
	data "github.com/Ocelani/swapi-planets/internal/database"
)

// MongoRepository is a repository of gen.Planet that uses the MongoDB.
type MongoRepository struct {
	db *data.MongoDB
}

// NewMongoRepository is a constructor of type NewMongoRepository repository.
func NewMongoRepository(mongoURI, collection string, planets []*gen.Planet) *MongoRepository {
	r := &MongoRepository{data.NewMongoDB(mongoURI, collection)}
	for _, pln := range planets {
		if err := r.Create(pln); err != nil {
			panic(err)
		}
	}
	return r
}

// Create just registers a gen.Planet data on database.
func (r *MongoRepository) Create(data *gen.Planet) error {
	mg, err := r.db.Collection.InsertOne(
		context.Background(),
		data,
	)
	if err != nil {
		return err
	}
	data.Id = mg.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

// ReadAll returns the entire data found on MongoRepository collection.
func (r *MongoRepository) ReadAll() ([]*gen.Planet, error) {
	cursor, err := r.db.Collection.Find(
		context.Background(),
		bson.M{},
	)
	if err != nil {
		return nil, err
	}
	planets := []*gen.Planet{}

	for cursor.Next(context.TODO()) {
		pln := &gen.Planet{}
		if err = cursor.Decode(&pln); err != nil {
			return nil, err
		}
		planets = append(planets, pln)
	}

	return planets, nil
}

// ReadOne finds and returns the data of a single gen.Planet with provided id.
func (r *MongoRepository) ReadOne(id string) (*gen.Planet, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	mg := r.db.Collection.FindOne(
		context.Background(),
		bson.M{"_id": oID},
	)
	p := &gen.Planet{}
	if err = mg.Decode(&p); err != nil {
		return nil, err
	}

	return p, nil
}

// Update searches the gen.Planet parameter ID, then, updates its data on database.
func (r *MongoRepository) Update(data *gen.Planet) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	var id struct {
		ID primitive.ObjectID `bson:"_id,omitempty"`
	}
	if err = json.Unmarshal(b, &id); err != nil {
		return err
	}
	_, err = r.db.Collection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": data},
	)
	if err != nil {
		return err
	}

	return nil
}

// Delete the specific gen.Planet data on database with its id as a parameter.
func (r *MongoRepository) Delete(id string) error {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.db.Collection.DeleteOne(
		context.Background(),
		bson.M{"_id": oID},
	)
	if err != nil {
		return err
	}

	return nil
}
