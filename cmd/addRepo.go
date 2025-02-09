package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// AppProject represents the ArgoCD AppProject YAML structure
type AppProject struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name       string   `yaml:"name"`
		Namespace  string   `yaml:"namespace"`
		Finalizers []string `yaml:"finalizers"`
	} `yaml:"metadata"`
	Spec struct {
		Description  string   `yaml:"description"`
		SourceRepos  []string `yaml:"sourceRepos"`
		Destinations []struct {
			Namespace string `yaml:"namespace"`
			Server    string `yaml:"server"`
			Name      string `yaml:"name"`
		} `yaml:"destinations"`
		ClusterResourceWhitelist []struct {
			Group string `yaml:"group"`
			Kind  string `yaml:"kind"`
		} `yaml:"clusterResourceWhitelist"`
		NamespaceResourceBlacklist []struct {
			Group string `yaml:"group"`
			Kind  string `yaml:"kind"`
		} `yaml:"namespaceResourceBlacklist"`
		NamespaceResourceWhitelist []struct {
			Group string `yaml:"group"`
			Kind  string `yaml:"kind"`
		} `yaml:"namespaceResourceWhitelist"`
		OrphanedResources struct {
			Warn bool `yaml:"warn"`
		} `yaml:"orphanedResources"`
		Roles []struct {
			Name        string   `yaml:"name"`
			Description string   `yaml:"description"`
			Policies    []string `yaml:"policies"`
			Groups      []string `yaml:"groups"`
			JwtTokens   []struct {
				Iat int64 `yaml:"iat"`
			} `yaml:"jwtTokens"`
		} `yaml:"roles"`
		SyncWindows []struct {
			Kind         string   `yaml:"kind"`
			Schedule     string   `yaml:"schedule"`
			Duration     string   `yaml:"duration"`
			Applications []string `yaml:"applications"`
			ManualSync   bool     `yaml:"manualSync"`
			Namespaces   []string `yaml:"namespaces"`
			Clusters     []string `yaml:"clusters"`
		} `yaml:"syncWindows"`
		PermitOnlyProjectScopedClusters bool     `yaml:"permitOnlyProjectScopedClusters"`
		SourceNamespaces                []string `yaml:"sourceNamespaces"`
	} `yaml:"spec"`
}

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
