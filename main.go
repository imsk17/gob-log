package main

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html"
	"log"
	"net/http"
)

// Embed the static/ directory to the app to serve favicons, custom css
// and all other `static` stuff.
//go:embed static
var static embed.FS

// Embed the Template Library
//go:embed templates/*
var template embed.FS

func main() {

	// Setup a Template Engine
	engine := html.NewFileSystem(http.FS(template), ".html")

	// Create a New Fiber App.
	// TODO: Use the `http` package instead.
	app := fiber.New(
		fiber.Config{
			Views: engine,
		})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("templates/index", fiber.Map{
			"Title": "again, mate.",
		})
	})

	// Mount the static directory
	app.Use("/", filesystem.New(
		filesystem.Config{
			Root:   http.FS(static),
			Browse: false,
		}))

	// Start the fiber app woo-hoo.
	log.Panic(app.Listen(":4000"))
}
