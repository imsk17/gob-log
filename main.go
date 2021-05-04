package main

import (
	"embed"
	"go-blog/src"
	"html/template"
	"log"
	"net/http"

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
//go:embed posts/*
var posts embed.FS

func main() {

	// Setup a Template Engine
	engine := html.NewFileSystem(http.FS(t), ".html")

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

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("templates/index", fiber.Map{
			"Title": "again, mate.",
		})
	})

	// Setup a Path for Blog
	app.Get("/:title", func(ctx *fiber.Ctx) error {
		title := ctx.Params("title", "")
		theme := ctx.Query("theme", "solarized-dark")
		if title == "" {
			return ctx.SendStatus(http.StatusNotFound)
		}
		blog, err := src.HTMLify(title, posts, theme)
		if err != nil {
			return ctx.SendStatus(http.StatusBadRequest)
		}
		return ctx.Render("templates/blog", blog)
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
