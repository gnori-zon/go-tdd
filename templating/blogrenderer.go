package blogrender

import (
	"embed"
	"github.com/gnori-zon/go-tdd/readingfiles/blogposts"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"html/template"
	"io"
	"strings"
)

var (
	//go:embed "templates/*"
	postTemplate embed.FS
)

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ}, nil
}

func (r *PostRenderer) Render(writer io.Writer, post blogposts.Post) error {
	postForRender := newPostData(post)
	return r.templ.ExecuteTemplate(writer, "blog.gohtml", postForRender)
}

func (r *PostRenderer) RenderIndex(writer io.Writer, posts []blogposts.Post) error {
	return r.templ.ExecuteTemplate(writer, "index.gohtml", newShortPostDataArray(posts))
}

type shortPostData struct {
	SanitizedTitle string
	Title          string
}

func newShortPostData(post blogposts.Post) shortPostData {
	return shortPostData{
		SanitizedTitle: newHref(post.Title),
		Title:          post.Title,
	}
}

func newShortPostDataArray(posts []blogposts.Post) any {
	shortPostDataArray := make([]shortPostData, 0, len(posts))
	for _, post := range posts {
		shortPostDataArray = append(shortPostDataArray, newShortPostData(post))
	}
	return shortPostDataArray
}

func newHref(title string) string {
	return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
}

type postData struct {
	title       string
	description string
	tags        []string
	body        template.HTML
}

func newPostData(post blogposts.Post) postData {
	return postData{
		title:       post.Title,
		description: post.Description,
		tags:        post.Tags,
		body:        template.HTML(mdToHTML(post.Body)),
	}
}

func mdToHTML(md string) string {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(md))

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(doc, renderer))
}
