package blogrender

import (
	"bytes"
	"github.com/approvals/go-approval-tests"
	"github.com/gnori-zon/go-tdd/readingfiles/blogposts"
	"io"
	"testing"
)

var post = blogposts.Post{
	Title:       "Testing Post",
	Description: "This is some description",
	Tags:        []string{"tag1", "tag2"},
	Body:        "# This is the **body**",
}

func TestRender(t *testing.T) {
	renderer, err := NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("post should convert post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		if err := renderer.Render(&buf, post); err != nil {
			t.Fatal("rendering failed", err)
		}
		approvals.VerifyString(t, buf.String())
	})

	t.Run("posts should convert index page HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogposts.Post{
			{Title: "First Post"},
			{Title: "Second Post"},
		}
		if err := renderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal("rendering failed", err)
		}
		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	renderer, err := NewPostRenderer()
	if err != nil {
		b.Fatal("rendering failed", err)
	}
	for b.Loop() {
		renderer.Render(io.Discard, post)
	}
}
