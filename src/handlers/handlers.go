package handlers

import (
	"embed"
	"go-blog/src/entities"
	"go-blog/src/parser"
	"log"
	"net/http"
	"net/url"
	"runtime"

	"github.com/gofiber/fiber/v2"
)

func blogHandler(mapBlogs *map[string]entities.Blog) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var title string
		if title, _ = url.PathUnescape(ctx.Params("title", "")); title == "" {
			return ctx.SendStatus(http.StatusNotFound)
		}
		blog, ok := (*mapBlogs)[title]
		if !ok {
			return ctx.SendStatus(http.StatusBadRequest)
		}
		return ctx.Render("templates/blog", fiber.Map{
			"Blog": blog,
			"Go":   runtime.Version(),
		})
	}
}

func blogsHandler(blogs *[]entities.Blog) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Render("templates/blogs", fiber.Map{
			"NoBlogs": false, // Use this switch to show hide all blogs.
			"Blogs":   blogs,
			"Go":      runtime.Version(),
		})
	}
}

func landingHandler(blogs *[]entities.Blog) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Render("templates/index", fiber.Map{
			"Go": runtime.Version(),
		})
	}
}

func SetupRoutes(app *fiber.App, posts embed.FS) {
	blogMap := map[string]entities.Blog{}
	var blogs, err = parser.ReadBlogs(posts)
	if err != nil {
		log.Fatal("Unable to parse blogs from the filesystem : ", err)
	}
	for _, v := range blogs {
		blogMap[v.Href] = v
	}
	app.Get("/blog/:title", blogHandler(&blogMap))
	app.Get("/blog", blogsHandler(&blogs))
	app.Get("/", landingHandler(&blogs))
}
