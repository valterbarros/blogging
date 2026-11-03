package content

import (
	"fmt"
	"os"
	"testing"
	"valterbarros/blog-tools/internal/utils"
)

func TestCreateDraft(t *testing.T) {
	CreateNewDraft("")

	path, _ := utils.GetFromRoot(draftsFolder)
	_, err := os.ReadFile(path + "/new-draft.md")

	if err != nil {
		t.Error("The file was not created")
	}
}

func TestConvertToPost(t *testing.T) {
	name := "draft-test.md"
	err := ConvertToPost(name)

	if err != nil {
		fmt.Println("an error occurred during convert", err)
	}
}
