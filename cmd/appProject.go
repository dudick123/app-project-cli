package cmd

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func appProjectRepoExists(appProject AppProject, repoURL string) bool {
	// Loop over the existing repositories
	for _, existingRepo := range appProject.Spec.SourceRepos {
		if existingRepo == repoURL {
			fmt.Println("Repository already exists in the AppProject YAML.")
			return true
		}
	}

	return false
}

func writeAppProjectFile(filePath string, appProject *AppProject) error {
	// Marshal the updated struct back to YAML
	updatedData, err := yaml.Marshal(&appProject)
	if err != nil {
		return fmt.Errorf("unable to convert updated structure to YAML: %v", err)
	}

	// Write back to the file
	err = os.WriteFile(filePath, updatedData, 0644)
	if err != nil {
		return fmt.Errorf("unable to write updated YAML to file: %v", err)
	}

	return nil
}

func openAppProjectFile(filePath string) (*AppProject, error) {
	// Read the AppProject YAML file
	data, err := os.ReadFile(filePath)
	if err != nil {
		err := fmt.Errorf("could not read the AppProject file at path: %v", filePath)
		return nil, err
	}

	// Unmarshal the YAML into AppProject struct
	var appProject AppProject
	err = yaml.Unmarshal(data, &appProject)
	if err != nil {
		err := fmt.Errorf("could not perform unmarshall on the the AppProject file path: %v", err)
		return nil, err
	}

	return &appProject, nil
}

func addRepoToAppProject(filePath, repoURL string) error {
	// Open the AppProject file
	appProject, err := openAppProjectFile(filePath)
	if err != nil {
		print("Error: %v\n", err)
		return err
	}

	// Check if the repo already exists
	if appProjectRepoExists(*appProject, repoURL) {
		return nil
	}

	// Add the repository
	appProject.Spec.SourceRepos = append(appProject.Spec.SourceRepos, repoURL)

	// Write the updated AppProject back to the file
	err = writeAppProjectFile(filePath, appProject)
	if err != nil {
		print("Error: %v\n", err)
		return err
	}

	fmt.Printf("Repository '%s' added to the AppProject YAML.\n", repoURL)
	return nil
}
