package main

import (
	"embed"
	"go-blog/src/handlers"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html"
)

// Embed the static/ directory to the app to serve favicons, custom css
// and all other `static` stuff.
//go:embed static
var static embed.FS

// Embed the Template Library
//go:embed templates/*
var t embed.FS

//Embed Our Actual Content aka The Blogs.
//go:embed posts/*.md
var posts embed.FS

func main() {

	PORT := "4000"
	OSPORT := os.Getenv("PORT")

	if OSPORT != "" {
		PORT = OSPORT
	} else {
		log.Printf("PORT Env Variable not set, Using the default port: %s Instead", PORT)
	}

	// Setup a Template Engine
	var engine *html.Engine

	if os.Getenv("GO_ENV") == "production" {
		engine = html.NewFileSystem(http.FS(t), ".html")
	} else {
		// For Dev Reasons
		engine = html.New(".", ".html")
		engine.Reload(true)
	}

	// We will be using this function to unescape the html which is coming from
	// goldmark markdown parser.
	engine.AddFunc(
		"unescape", func(s string) template.HTML {
			return template.HTML(s)
		},
	)

	// Create a New Fiber App.
	// TODO: Use the `http` package instead.
	app := fiber.New(
		fiber.Config{
			Views: engine,
		})

	// Setup handlers for the blogs
	handlers.SetupRoutes(app, posts)

	// Mount the static directory
	if os.Getenv("GO_ENV") == "production" {
		app.Use("/", filesystem.New(
			filesystem.Config{
				Root:   http.FS(static),
				Browse: false,
			}))
	} else {
		app.Static("/static", "static/")
	}

	// Start the fiber app woo-hoo.
	log.Panic(app.Listen(":" + PORT))
}
