package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func Render(w io.Writer, p Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.html")
	if err != nil {
		return err
	}

	if err := templ.ExecuteTemplate(w,"blog.html", p); err != nil {
		return err
	}

	return nil
}
