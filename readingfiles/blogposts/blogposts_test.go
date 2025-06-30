package blogposts

import (
	"errors"
	"io/fs"
	"reflect"
	"strings"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct{}

func (fs StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("fail")
}

type TestFile struct {
	name     string
	wantPost Post
}

func toFs(files []TestFile) fstest.MapFS {
	fstTemp := fstest.MapFS{}
	for _, file := range files {
		fstTemp[file.name] = &fstest.MapFile{
			Data: []byte("Title: " + file.wantPost.Title +
				"\nDescription: " + file.wantPost.Description +
				"\nTags: " + strings.Join(file.wantPost.Tags, ",") +
				"\n---" +
				"\n" + file.wantPost.Body),
		}
	}
	return fstTemp
}

func TestNewBlogPosts(t *testing.T) {

	t.Run("success read posts", func(t *testing.T) {
		cases := []struct {
			files []TestFile
		}{
			{
				files: []TestFile{
					{
						name: "post_1.md",
						wantPost: Post{
							Title:       "Post 1",
							Description: "Description 1",
							Tags:        []string{"Tag 1", "Tag 2"},
							Body:        "Some body for post 1\nDescription is correct and hello",
						},
					},
					{
						name: "post_2.md",
						wantPost: Post{
							Title:       "Post 2",
							Description: "Description 2",
							Tags:        []string{"Tag 3", "Tag 4"},
							Body:        "Some body for post 2\nDescription is correct and halo",
						},
					},
				},
			},
		}
		for _, testCase := range cases {
			fileSystem := toFs(testCase.files)
			posts, err := NewPostsFromFS(fileSystem)

			if err != nil {
				t.Fatal(err)
			}

			if len(posts) != len(fileSystem) {
				t.Errorf("want %d length of posts but got %d", len(fileSystem), len(posts))
			}

			for _, testFile := range testCase.files {
				if !containsPost(posts, testFile.wantPost) {
					t.Errorf("want posts contains post %+v but got %+v", testFile.wantPost, posts)
				}
			}
		}
	})

	t.Run("error read posts", func(t *testing.T) {
		fileSystem := StubFailingFS{}
		_, err := NewPostsFromFS(fileSystem)

		if err == nil {
			t.Error("want error but got nil")
		}
	})
}

func containsPost(posts []Post, wantPost Post) bool {
	for _, post := range posts {
		if reflect.DeepEqual(post, wantPost) {
			return true
		}
	}
	return false
}
