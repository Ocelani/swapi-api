package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/Ocelani/swapi-planets/gen"
	"github.com/Ocelani/swapi-planets/internal/swapi"
	"github.com/Ocelani/swapi-planets/pkg/planet"
)

const (
	// _port is the default port of this API.
	_port = 80

	// _mongoURL is the default url that connects to MongoDB.
	_mongoURL = "mongodb://admin:admin@mongo:27017/"
)

// _log is a private instance of this API logging instance.
var _log zerolog.Logger = Logger()

// Logger starts the logger of this API.
func Logger() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	return log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
}

// Accepts is the application header accepted content type.
func Accepts(c *fiber.Ctx) error {
	c.Accepts("application/json", "text/html", "html", "text", "json")
	return nil
}

// Cors stands for Cross Origins Rources Sharing.
func Cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, X-Auth",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowOrigins:     "*",
		AllowCredentials: true,
	})
}

// main entrypoint.
func main() {
	port := flag.Int("port", _port, "exposed port of this API")
	skipData := flag.Bool("skip-data", false, "skip getting the provided json data")
	cacheDB := flag.Bool("cache-db", false, "use cacheDB adapter (default MongoDB)")
	flag.Parse()

	planets := []*gen.Planet{}
	if !*skipData {
		planets = swapi.GetPlanets()
	}

	var repo planet.Repository
	if *cacheDB {
		repo = planet.NewCacheRepository(planets)
	} else {
		repo = planet.NewMongoRepository(_mongoURL, "planets", planets)
	}

	a := &API{
		App:     fiber.New(),
		Service: planet.NewDefaultService(repo),
	}

	defer func() {
		if err := a.App.Shutdown(); err != nil {
			_log.Panic().Err(err).Send()
		}
	}()

	a.App.Use(Cors())
	a.Router()

	if err := a.App.Listen(fmt.Sprintf(":%d", *port)); err != nil {
		log.Panic().Err(err)
	}
}
