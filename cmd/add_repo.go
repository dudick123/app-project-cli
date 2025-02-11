package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// addRepoCmd represents the add-repo command
var addRepoCmd = &cobra.Command{
	Use:   "add-repo",
	Short: "Add a repository to an ArgoCD AppProject YAML file if it does not exist",
	Run: func(cmd *cobra.Command, args []string) {
		filePath, _ := cmd.Flags().GetString("file")
		repoURL, _ := cmd.Flags().GetString("repo")

		if filePath == "" || repoURL == "" {
			fmt.Println("Error: --file and --repo flags are required")
			os.Exit(1)
		}

		err := addRepoToAppProject(filePath, repoURL)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		//fmt.Println("Repository successfully added (if it was missing).")
	},
}

func init() {
	rootCmd.AddCommand(addRepoCmd)
	addRepoCmd.Flags().StringP("file", "f", "", "Path to the AppProject YAML file")
	addRepoCmd.Flags().StringP("repo", "r", "", "Repository URL to add")
	addRepoCmd.MarkFlagRequired("file")
	addRepoCmd.MarkFlagRequired("repo")
}
