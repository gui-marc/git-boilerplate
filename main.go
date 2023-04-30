package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/gui-marc/git-boilerplate/utils"
)

func main() {
	utils.Clear()

	if len(os.Args) < 3 {
		utils.Error("Error: invalid arguments\n")
		utils.Info("Usage: git-boilerplate <template-repo> <project-name>")
		return
	}

	templateRepo := os.Args[1]
	projectName := os.Args[2]

	utils.Loading("Cloning template repository...\n")
	output, err := exec.Command("git", "clone", "--depth=1", "--branch=main", templateRepo, ".git-template").CombinedOutput()
	if err != nil {
		utils.Error(fmt.Sprintf("Error cloning template repository: %v\n%s", err, output))
		return
	}

	utils.Info(fmt.Sprintf("Creating project directory %s...\n", projectName))
	err = os.Mkdir(projectName, os.ModePerm)
	if err != nil {
		utils.Error(fmt.Sprintf("Error creating project directory: %v\n", err))
		return
	}

	utils.Info("Copying template contents to project directory...")
	copyDir(".git-template", projectName)

	utils.Info("Initializing new Git repository...")
	err = os.RemoveAll(fmt.Sprintf("%s/.git", projectName))
	if err != nil {
		utils.Error(fmt.Sprintf("Error removing .git folder: %v\n", err))
		return
	}

	_, err = exec.Command("git", "init", projectName).CombinedOutput()
	if err != nil {
		utils.Error(fmt.Sprintf("Error initializing new Git repository: %v\n", err))
		return
	}

	utils.Info("Removing template directory...")
	err = os.RemoveAll(".git-template")
	if err != nil {
		utils.Error(fmt.Sprintf("Error removing template directory: %v\n", err))
		return
	}

	utils.Success("Done!")
}

func copyDir(src string, dest string) error {
	// Read source directory
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	// Create destination directory
	err = os.Mkdir(dest, os.ModePerm)
	if err != nil {
		return err
	}

	// Copy files
	for _, file := range files {
		srcFile := src + "/" + file.Name()
		destFile := dest + "/" + file.Name()

		if file.IsDir() {
			err = copyDir(srcFile, destFile)
			if err != nil {
				return err
			}
		} else {
			content, err := ioutil.ReadFile(srcFile)
			if err != nil {
				return err
			}

			// Replace template variables with project-specific values
			contentStr := strings.Replace(string(content), "{{project-name}}", dest, -1)

			err = ioutil.WriteFile(destFile, []byte(contentStr), file.Mode())
			if err != nil {
				return err
			}
		}
	}

	return nil
}
