package content

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"valterbarros/blog-tools/internal/utils"
)

type Manage struct {
	folder string
}

func NewManage(folder string) *Manage {
	return &Manage{folder}
}

func (m Manage) ListSome() (string, error) {
	path, err := utils.GetFromRoot(m.folder)
	files, err := utils.ListFiles(path)

	if err != nil {
		return "", err
	}

	return strings.Join(files, "\n"), nil
}

func (m Manage) CreateSome(name string) error {
	if name == "" || filepath.Ext(name) == "" {
		name = name + ".md"
	}

	path, _ := utils.GetFromRoot(m.folder)

	err := utils.SaveFile(filepath.Join(path, name), "")

	if err != nil {
		fmt.Println("The file was not created: ", err)
		return err
	}

	return nil
}

func (m Manage) DeleteSome(name string) error {
	path, _ := utils.GetFromRoot(m.folder)

	err := utils.RemoveFile(filepath.Join(path, name))

	if err != nil {
		fmt.Println("The file was not removed: ", err)
		return err
	}

	return nil
}

func (m Manage) MoveSome(name, newPath string) error {
	if name == "" || filepath.Ext(name) == "" {
		name = name + ".md"
	}

	pathRoot, _ := utils.GetFromRoot(m.folder)
	newRoot, _ := utils.GetFromRoot(newPath)

	oldPathFile := filepath.Join(pathRoot, name)
	newPathFile := filepath.Join(newRoot, name)

	fmt.Println(oldPathFile, newPathFile)

	return os.Rename(oldPathFile, newPathFile)
}
