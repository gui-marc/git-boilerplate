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
		utils.Error("Invalid arguments")
		utils.Info("Usage: git-boilerplate <template-url> <project-name>")
		return
	}

	templateURL := os.Args[1]
	projectName := os.Args[2]

	utils.Loading("Cloning template repository...")
	output, err := exec.Command("git", "clone", "--depth=1", "--branch=main", templateURL, ".git-boilerplate-temp").CombinedOutput()

	if err != nil {
		utils.Error("Error cloning template repository")
		fmt.Printf("%v\n%s\n", err, string(output))
		return
	}

	utils.Success("Template repository cloned successfully")
	utils.Loading("Removing .git directory...")

	err = os.RemoveAll(".git-boilerplate-temp/.git")

	if err != nil {
		utils.Error("Error removing .git directory")
		return
	}

	utils.Success(".git directory removed successfully")

	utils.Loading("Copying to new folder...")

	err = copyDir(".git-boilerplate-temp", projectName, projectName)

	if err != nil {
		utils.Error("Error copying files")
		return
	}

	utils.Success("Files copied successfully")

	utils.Loading("Removing temporary directory...")

	err = os.RemoveAll(".git-boilerplate-temp")

	if err != nil {
		utils.Error("Error removing temporary directory")
		return
	}

	utils.Success("Temporary directory removed successfully")

	utils.Loading("Initializing git repository...")

	output, err = exec.Command("git", "init", projectName).CombinedOutput()

	if err != nil {
		utils.Error("Error initializing git repository")
		fmt.Printf("%v\n%s\n", err, string(output))
		return
	}

	utils.Success("Git repository initialized successfully")

	utils.Clear()

	utils.Success("Project created successfully 🎉\n")
	utils.Info("To start working on your project, run the following command:")
	utils.Info(fmt.Sprintf("cd %s", projectName))
}

func copyDir(src string, dest string, projectName string) error {
	// Read source directory
	files, err := ioutil.ReadDir(fmt.Sprintf("./%s", src))
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
			err = copyDir(srcFile, destFile, projectName)
			if err != nil {
				return err
			}
		} else {
			content, err := ioutil.ReadFile(srcFile)
			if err != nil {
				return err
			}

			// Replace template variables with project-specific values
			contentStr := strings.Replace(string(content), "{{project-name}}", projectName, -1)

			err = ioutil.WriteFile(destFile, []byte(contentStr), file.Mode())
			if err != nil {
				return err
			}
		}
	}

	return nil
}
