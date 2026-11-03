package content

var draftsFolder = "articles/drafts"

var mnge = NewManage(draftsFolder)

func ListDrafts() (string, error) {
	return mnge.ListSome()
}

func CreateNewDraft(name string) error {
	return mnge.CreateSome(name)
}

func DeleteDraft(name string) error {
	return mnge.DeleteSome(name)
}

func ConvertToPost(name string) error {
	return mnge.MoveSome(name, postsFolder)
}
