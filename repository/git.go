package repository

import (
	"io/ioutil"
	"os"
	"path"

	// Third-Party
	"github.com/go-git/go-git"

	// Local
	"github.com/alvesgabriel/cookiecutter/utils"
)

// Init inicalizes git if directory doesn't have .git directory
func Init(directory string) {
	if _, err := os.Stat(path.Join(directory, ".git")); os.IsNotExist(err) {
		_, err := git.PlainInit(directory, false)
		utils.FatalError(err)
		createGitignore(directory)
	}
}

func createGitignore(directory string) {
	gitignorePath := path.Join(directory, ".gitignore")
	if _, err := os.Stat(gitignorePath); os.IsNotExist(err) {
		err := ioutil.WriteFile(gitignorePath, gitignore, 0664)
		utils.FatalError(err)
	}
}
