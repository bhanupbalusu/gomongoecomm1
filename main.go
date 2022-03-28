package main

import (
	u "github.com/bhanupbalusu/gomongoecomm1/utils"
	"github.com/gofiber/fiber/v2" // add Fiber package
)

func main() {
	app := fiber.New() // create a new Fiber instance

	h := u.Init()
	// Create a new endpoint
	app.Get("/products", h.Get)
	app.Get("/product/:id", h.GetByID)
	app.Post("/product", h.Create)
	app.Put("/product/:id", h.Update)
	app.Delete("/product/:id", h.Delete)

	// Start server on port 3000
	app.Listen(":3000")
}
