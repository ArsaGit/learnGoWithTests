package blogrenderer_test

import (
	"io"

	approvals "github.com/approvals/go-approval-tests"

	"bytes"
	"learngowithtests/hello/blogrenderer"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := blogrenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is description",
			Tags:        []string{"go", "tdd"},
		}
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blogrenderer.Render(io.Discard, aPost)
	}
}
