package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

type PostViewModel struct {
	Title, SanitisedTitle, Description, Body string
	Tags                                     []string
}

type PostRenderer struct {
	templ    *template.Template
	mdParser *parser.Parser
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

func Render(w io.Writer, p Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.html")
	if err != nil {
		return err
	}

	if err := templ.ExecuteTemplate(w, "blog.html", p); err != nil {
		return err
	}

	return nil
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.html")
	if err != nil {
		return nil, err
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	return &PostRenderer{templ: templ, mdParser: parser}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.html", newPostVM(p, r))
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.html", posts)
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

func newPostVM(p Post, r *PostRenderer) postViewModel {
	vm := postViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return vm
}
