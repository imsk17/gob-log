package src

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"go-blog/src/entities"
	"go-blog/src/errs"
	"strings"
	"time"
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
		Summary:     metaData["Summary"].(string),
		PublishDate: metaData["PublishDate"].(string),
		SkimTime:    metaData["SkimTime"].(string),
		Author:      metaData["Author"].(string),
	}
	return b, nil
}

func ReadBlogs(f embed.FS) ([]entities.BlogMeta, error) {
	var blogs []entities.BlogMeta
	allBlogs, _ := f.ReadDir("posts")
	for _, v := range allBlogs {
		splits := strings.Split(v.Name(), "-")
		t, _ := time.Parse("02012006", splits[1])
		b := entities.BlogMeta{
			SkimTime: fmt.Sprintf("%v Mins Read", splits[0]),
			Title: strings.ReplaceAll(strings.ReplaceAll(splits[2], "_", " "), ".md", ""),
			PublishDate: strings.ReplaceAll(t.Format(time.RFC1123), "00:00:00 UTC", ""),
			Href: strings.ReplaceAll(v.Name(), ".md" ,""),
			Tags: strings.Split(splits[3][1:len(splits[3])-4], " "),
		}
		blogs = append(blogs, b)
	}
	return blogs, nil
}
