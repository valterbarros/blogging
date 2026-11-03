package content

import (
	"fmt"
	"strings"
	"testing"
	"valterbarros/blog-tools/internal/utils"
)

func TestListSome(t *testing.T) {
	dir := t.TempDir()
	mn := NewManage(dir)
	createFiles(t, dir)
	files, err := mn.ListSome()

	if err != nil {
		fmt.Println("an error occured", err)
	}

	total := len(strings.Split(files, "\n"))

	if total != 3 {
		t.Error("Just found this files: ", total)
	}
}

func TestCreateSome(t *testing.T) {
	dir := t.TempDir()
	mn := NewManage(dir)
	name := "new-post.md"
	err := mn.CreateSome(name)

	if err != nil {
		fmt.Println("an error occured", err)
	}

	files, err := mn.ListSome()
	total := len(strings.Split(files, "\n"))

	if total != 1 {
		t.Error("Just found this files: ", total)
	}
}

func TestDeleteSome(t *testing.T) {
	dir := t.TempDir()
	mn := NewManage(dir)
	createFiles(t, dir)

	err := mn.DeleteSome("test3.md")

	if err != nil {
		fmt.Println("an error occured", err)
	}

	files, err := mn.ListSome()
	total := len(strings.Split(files, "\n"))

	if total != 2 {
		t.Error("Just found this files: ", total)
	}
}

func TestMoveSome(t *testing.T) {
	dir := t.TempDir()
	mn := NewManage(dir)
	createFiles(t, dir)

	dir2 := t.TempDir()

	err := mn.MoveSome("test3.md", dir2)

	if err != nil {
		fmt.Println("an error occured", err)
	}

	files, _ := utils.ListFiles(dir2)

	if len(files) != 1 {
		t.Error("Just found this files: ", files)
	}
}

func createFiles(t *testing.T, dir string) {
	t.Helper()

	err := utils.SaveFile(dir+"/test1.md", "123")
	if err != nil {
		fmt.Println("an error occured", err)
	}
	utils.SaveFile(dir+"/test3.md", "123")
	utils.SaveFile(dir+"/test4.md", "123")
}
