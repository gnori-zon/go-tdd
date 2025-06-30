package blogposts

import (
	"bufio"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titlePrefix       = "Title: "
	descriptionPrefix = "Description: "
	tagsPrefix        = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := newPostScanner(postFile)
	return Post{
		Title:       scanner.readMetaLine(titlePrefix),
		Description: scanner.readMetaLine(descriptionPrefix),
		Tags:        trimSpace(strings.Split(scanner.readMetaLine(tagsPrefix), ",")),
		Body:        scanner.readBody(),
	}, nil
}

type postScanner struct {
	*bufio.Scanner
}

func newPostScanner(reader io.Reader) *postScanner {
	return &postScanner{bufio.NewScanner(reader)}
}

func (s *postScanner) readMetaLine(tagName string) string {
	s.Scan()
	return strings.TrimSpace(strings.TrimPrefix(s.Text(), tagName))
}

func (s *postScanner) readBody() string {
	s.Scan()
	var stringBuilder strings.Builder
	for s.Scan() {
		stringBuilder.WriteString(s.Text())
		stringBuilder.WriteString("\n")
	}
	return strings.TrimSuffix(stringBuilder.String(), "\n")
}

func trimSpace(items []string) []string {
	result := make([]string, 0, len(items))
	for _, item := range items {
		result = append(result, strings.TrimSpace(item))
	}
	return result
}
