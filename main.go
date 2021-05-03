package main

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"log"
	"net/http"
)

// Embed the static/ directory to the app to serve favicons, custom css
// and all other `static` stuff.
//go:embed static
var static embed.FS

func main() {
	// Create a New Fiber App.
	// TODO: Use the `http` package instead.
	app := fiber.New()

	// Mount the static directory
	app.Use("/", filesystem.New(
		filesystem.Config{
			Root:   http.FS(static),
			Browse: false,
		}))

	// Start the fiber app woo-hoo.
	log.Panic(app.Listen(":4000"))
}
