package main

import (
	"fmt"
	"log"
	"strings"
	"valterbarros/blog-tools/internal/content"
	"valterbarros/blog-tools/internal/editor"
	"valterbarros/blog-tools/internal/utils"

	"github.com/spf13/cobra"
)

func draftCommands() *cobra.Command {
	draftsCmd := &cobra.Command{
		Use:   "drafts",
		Short: "Manage drafts.",
		Args:  cobra.NoArgs,
	}

	listDraftCmd := &cobra.Command{
		Use:   "list",
		Short: "List drafts",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			drafts, err := content.ListDrafts()
			if err != nil {
				return fmt.Errorf("list drafts: %w", err)
			}
			fmt.Println(drafts)
			return nil
		},
	}

	deleteDraftCmd := &cobra.Command{
		Use:   "delete <name>",
		Short: "Delete a draft",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := content.DeleteDraft(args[0]); err != nil {
				return fmt.Errorf("delete draft: %w", err)
			}
			return nil
		},
	}

	convertDraftCmd := &cobra.Command{
		Use:   "convert <name>",
		Short: "Convert a draft to post",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := content.ConvertToPost(args[0]); err != nil {
				return fmt.Errorf("convert draft: %w", err)
			}
			return nil
		},
	}

	createDraftCmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create a draft",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := content.CreateNewDraft(args[0]); err != nil {
				return fmt.Errorf("create draft: %w", err)
			}
			return nil
		},
	}

	draftsCmd.AddCommand(listDraftCmd, deleteDraftCmd, convertDraftCmd, createDraftCmd)
	return draftsCmd
}

func postsCommands() *cobra.Command {
	postCmd := &cobra.Command{
		Use:   "posts",
		Short: "utils to create/list/delete posts.",
		Args:  cobra.ArbitraryArgs,
	}

	listPostCmd := &cobra.Command{
		Use:   "list",
		Short: "List posts",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			posts, err := content.ListPosts()
			if err != nil {
				return fmt.Errorf("list posts: %w", err)
			}
			fmt.Println(posts)
			return nil
		},
	}

	createPostCmd := &cobra.Command{
		Use:   "create",
		Short: "Create post",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := content.CreateNewPost(args[0])
			if err != nil {
				return fmt.Errorf("create posts: %w", err)
			}
			return nil
		},
	}

	convertPostCmd := &cobra.Command{
		Use:   "convert",
		Short: "Convert post to draft",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := content.ConvertToDraft(args[0])
			if err != nil {
				return fmt.Errorf("convert posts: %w", err)
			}
			return nil
		},
	}

	deletePostCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a post",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := content.DeletePost(args[0])
			if err != nil {
				return fmt.Errorf("delete post: %w", err)
			}
			return nil
		},
	}

	formatPostCmd := &cobra.Command{
		Use:   "format",
		Short: "Format a post",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := content.FormatPost(args[0], "")
			if err != nil {
				return fmt.Errorf("format post: %w", err)
			}
			return nil
		},
	}

	editorPostCmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit posts",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := editor.RunServer()
			if err != nil {
				return fmt.Errorf("run editor: %w", err)
			}
			return nil
		},
	}

	postCmd.AddCommand(listPostCmd, createPostCmd, convertPostCmd, deletePostCmd, formatPostCmd, editorPostCmd)

	return postCmd
}

func commitCommands() *cobra.Command {
	return &cobra.Command{
		Use:   "commit",
		Short: "commit data.",
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			changes := utils.GetChanges()

			if strings.Contains(changes, "articles") {
				path, _ := utils.GetFromRoot("articles")
				utils.CommitChanges("[cli blogging] articles updated", path)
			} else {
				fmt.Println("not staged articles/**/* folder")
			}
		},
	}
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "blogging",
		Short: "Manage blog CLI.",
	}

	rootCmd.AddCommand(draftCommands())
	rootCmd.AddCommand(postsCommands())
	rootCmd.AddCommand(commitCommands())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
