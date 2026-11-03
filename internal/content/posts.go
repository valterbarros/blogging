package content

import (
	"fmt"
	"path/filepath"
	"strings"
	"valterbarros/blog-tools/internal/utils"
)

var postsFolder = "articles/_posts"

var mn = NewManage(postsFolder)

func ListPosts() (string, error) {
	return mn.ListSome()
}

func CreateNewPost(name string) error {
	return mn.CreateSome(name)
}

func DeletePost(name string) error {
	return mn.DeleteSome(name)
}

func ConvertToDraft(name string) error {
	return mn.MoveSome(name, draftsFolder)
}

// specific methods
func FormatPost(name string, folder string) error {
	if folder == "" {
		folder = postsFolder
	}

	postsPath, _ := utils.GetFromRoot(folder)
	currentPostFile := filepath.Join(postsPath, name)
	data, _ := utils.GetFileData(currentPostFile)

	metadata := addMetadata(data)
	err := utils.SaveFile(currentPostFile, metadata+data)
	if err != nil {
		fmt.Println("error: ", err)
	}

	newName := addDateToName(postsPath, name)

	return utils.RenameFile(currentPostFile, newName)
}

func addDateToName(postsPath, name string) string {
	now, _ := utils.GetDate()

	renamedPostFile := filepath.Join(postsPath, now+"-"+name)
	return renamedPostFile
}

func addMetadata(justified string) string {
	if strings.Contains(justified, "layout:") {
		return ""
	}
	metadata := `---
layout: post
title: [title]
---
`

	return metadata
}
