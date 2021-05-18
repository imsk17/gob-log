package src

import (
	"bytes"
	"embed"
	"fmt"
	"go-blog/src/entities"
	"go-blog/src/errs"

	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	highlighting "github.com/yuin/goldmark-highlighting"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// This function is the whole and soul of this web app.
// It parses the html from the markdown file and converts
// into HTML. Also, Extracts metadata about the blog.
func HTMLify(filename string, fs embed.FS, theme string) (entities.Blog, error) {
	path := fmt.Sprintf("posts/%s", filename)
	var b entities.Blog
	markdown := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle(theme),
				highlighting.WithFormatOptions(),
			),
			meta.Meta,
			emoji.Emoji,
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)
	context := parser.NewContext()
	blog, err := fs.ReadFile(path)
	if err != nil {
		return b, errs.ErrNoBlogFound
	}
	var buf bytes.Buffer
	if err := markdown.Convert(blog, &buf, parser.WithContext(context)); err != nil {
		panic(err)
	}
	metaData := meta.Get(context)
	idx, _ := metaData["Index"].(int)
	skim, _ := metaData["SkimTime"].(int)
	b = entities.Blog{
		Index:    idx,
		Title:    metaData["Title"].(string),
		Dept:     metaData["Dept"].(string),
		Content:  buf.String(),
		Subtopic: metaData["Subtopic"].(string),
		Date:     metaData["Date"].(string),
		Time:     metaData["Time"].(string),
		SkimTime: skim,
		Author:   metaData["Author"].(string),
		Topic:    metaData["Topic"].(string),
		Href:     filename,
	}
	return b, nil
}

func ReadBlogs(f embed.FS) ([]entities.Blog, error) {
	var blogs []entities.Blog
	allBlogs, _ := f.ReadDir("posts")
	for _, v := range allBlogs {
		b, _ := HTMLify(v.Name(), f, "dracula")
		blogs = append(blogs, b)
	}
	return blogs, nil
}
