package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ListFiles(path string) ([]string, error) {
	files, err := os.ReadDir(path)

	if err != nil {
		return []string{}, err
	}

	names := []string{}

	for _, f := range files {
		names = append(names, f.Name())
	}

	return names, nil
}

const repoName = "valterbarros.github.io"

// TODO: add scenarios about use path with ../
func GetFromRoot(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	}

	pwd, _ := os.Getwd()

	if !strings.Contains(pwd, repoName) {
		return "", fmt.Errorf("go to right path")
	}

	base, _, _ := strings.Cut(pwd, repoName)

	return filepath.Join(base, repoName, path), nil
}

func SaveFile(file, data string) error {
	return os.WriteFile(file, []byte(data), 0644)
}

func RemoveFile(file string) error {
	return os.Remove(file)
}

func RenameFile(orig, dest string) error {
	return os.Rename(orig, dest)
}

func ResetFile(file string) {
	os.WriteFile(file, []byte(""), 0644)
}

func GetFileData(file string) (string, error) {
	data, err := os.ReadFile(file)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
