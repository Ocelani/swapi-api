package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/Ocelani/swapi-planets/gen"
	"github.com/Ocelani/swapi-planets/pkg/planet"
)

// Response type represents a common API response.
type Response struct {
	Data  interface{}
	Error error
}

// API stores this application resources.
type API struct {
	App     *fiber.App
	Service planet.Service
}

// Router defines the URL routing properties.
func (a *API) Router() {
	a.App.Get("/", Accepts)
	a.App.Post("/planet", a.create)
	a.App.Get("/planet", a.readAll)
	a.App.Get("/planet/:id", a.readOne)
	a.App.Put("/planet", a.update)
	a.App.Delete("/planet/:id", a.delete)
}

// add through API service request.
func (a *API) create(c *fiber.Ctx) error {
	pln := &gen.Planet{}

	// Parses request body pln data
	if err := c.BodyParser(pln); err != nil {
		_log.Err(err).Msg("on planet.create body parser")
		return c.Status(http.StatusBadRequest).JSON(Response{
			Data:  pln,
			Error: err,
		})
	}
	_log.Info().
		Str("id", pln.Id).
		Str("name", pln.Name).
		Str("climate", pln.Climate).
		Str("terrain", pln.Terrain).
		Msg("create pln")

	// Creates and saves this pln data
	if err := a.Service.Create(pln); err != nil {
		_log.Err(err).Send()
		return c.Status(http.StatusUnprocessableEntity).JSON(Response{
			Data:  pln,
			Error: err,
		})
	}

	// Send data response
	return c.Status(http.StatusOK).JSON(
		Response{Data: pln},
	)
}

// getAll through API client request.
func (a *API) readAll(c *fiber.Ctx) error {
	// Get all saved plns
	result, err := a.Service.ReadAll()

	if err != nil {
		_log.Err(err).Msg("read all planets")
		return c.Status(http.StatusUnprocessableEntity).JSON(Response{
			Data:  result,
			Error: err,
		})
	}
	_log.Info().Msg("readed all planets")

	// Send data response
	return c.Status(http.StatusOK).JSON(
		Response{Data: result},
	)
}

// getOneWithID through API client request.
func (a *API) readOne(c *fiber.Ctx) error {
	id := c.Params("id")

	// Get this pln with id
	result, err := a.Service.ReadOne(id)
	if err != nil {
		_log.Err(err).Msg("read one planet")
		return c.Status(http.StatusUnprocessableEntity).JSON(Response{
			Data:  id,
			Error: err,
		})
	}
	_log.Info().
		Str("id", id).
		Str("name", result.Name).
		Str("climate", result.Climate).
		Str("terrain", result.Terrain).
		Msg("readed one planet")

	// Send data response
	return c.Status(http.StatusOK).JSON(
		Response{Data: result},
	)
}

// update through API service request.
func (a *API) update(c *fiber.Ctx) error {
	pln := &gen.Planet{}

	// Parses request body planet data
	if err := c.BodyParser(&pln); err != nil {
		_log.Err(err).Msg("on planet.update body parser")
		return c.Status(400).JSON(Response{
			Data:  pln,
			Error: err,
		})
	}

	if err := a.Service.Update(pln); err != nil {
		_log.Err(err).Msg("planet update")
		return c.Status(400).JSON(Response{
			Data:  pln,
			Error: err,
		})
	}
	_log.Info().
		Str("id", pln.Id).
		Msg("planet updated")

	return c.Status(200).JSON(Response{
		Data:  pln,
		Error: nil,
	})
}

// remove through API client request.
func (a *API) delete(c *fiber.Ctx) error {
	id := c.Params("id")

	// Delete this id
	if err := a.Service.Delete(id); err != nil {
		_log.Err(err).Msg("planet delete")
		return c.Status(http.StatusUnprocessableEntity).JSON(Response{
			Data:  id,
			Error: err,
		})
	}
	_log.Info().
		Str("id", id).
		Msg("planet deleted")

	// Send data response
	return c.Status(http.StatusOK).JSON(
		Response{Data: id},
	)
}
