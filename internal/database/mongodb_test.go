package database

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

var tDB = struct {
	url  string
	coll string
}{
	"mongodb://admin:admin@localhost:27017/",
	"tests",
}

var mongoT = flag.Bool("mongo", false, "runs mongo-db tests")

func TestDatabase_NewMongoDB(t *testing.T) {
	flag.Parse()
	if !*mongoT {
		return
	}
	var want *MongoDB
	got := NewMongoDB(tDB.url, tDB.coll)
	assert.IsType(t, want, got)
}

func TestDatabase_ConnectMongoDB(t *testing.T) {
	flag.Parse()
	if !*mongoT {
		return
	}
	var want *mongo.Database
	got := ConnectMongoDB(tDB.url)
	assert.IsType(t, want, got)
	assert.Equal(t, "sigma", got.Name())
}
