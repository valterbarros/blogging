package utils

import (
	"strings"
	"testing"
)

func TestFiles(t *testing.T) {
	path, _ := GetFromRoot("articles/_posts")
	files, _ := ListFiles(path)

	if len(files) < 1 {
		t.Error("The directory is empty")
	}
}

func TestGoFromRoot(t *testing.T) {
	t.Run("should return from whatever path", func(t *testing.T) {
		path, _ := GetFromRoot("new_path/other")

		if !strings.Contains(path, "new_path/other") {
			t.Error("The path is wrong")
		}
	})

	t.Run("should return just project root path", func(t *testing.T) {
		path, _ := GetFromRoot(".")

		end := strings.HasSuffix(path, "valterbarros.github.io")

		if !end {
			t.Error("The path is not root project")
		}
	})

	t.Run("should return same path when is abs", func(t *testing.T) {
		path := "/tmp/file1.txt"
		root, _ := GetFromRoot("/tmp/file1.txt")

		if root != path {
			t.Error("root shoud be equal")
		}
	})
}
