package src

import (
	"bytes"
	"embed"
	"fmt"
	"go-blog/src/entities"
	"go-blog/src/errs"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

// This function is the whole and soul of this web app.
// It parses the html from the markdown file and converts
// into HTML. Also, Extracts metadata about the blog.
func HTMLify(filename string, fs embed.FS, theme string) (entities.Blog, error) {
	path := fmt.Sprintf("posts/%s.md", filename)
	var b entities.Blog
	markdown := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle(theme),
				highlighting.WithFormatOptions(),
			),
			meta.Meta,
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
	b = entities.Blog{
		Title:       metaData["Title"].(string),
		Content:     buf.String(),
		Tags:        metaData["Tags"].([]interface{}),
		Image:       metaData["Image"].(string),
		Summary:     metaData["Summary"].(string),
		PublishDate: metaData["PublishDate"].(string),
		SkimTime:    metaData["SkimTime"].(string),
	}
	return b, nil
}
