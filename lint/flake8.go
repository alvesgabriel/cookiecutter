package lint

import (
	"fmt"
	"log"
	"os"
	"path"

	// Local
	"github.com/alvesgabriel/cookiecutter/packages"
	"github.com/alvesgabriel/cookiecutter/utils"
)

// CreateFile create file '.flake8'
func CreateFile(directory string) {
	flake8Content := `[flake8]
	max-line-length=120
	exclude=.venv`
	flake8File := path.Join(directory, ".flake8")

	if _, err := os.Stat(flake8File); os.IsNotExist(err) {
		f, err := os.OpenFile(flake8File, os.O_CREATE, 0664)
		utils.FatalError(err)
		defer f.Close()

		if _, err := f.Write([]byte(flake8Content)); err != nil {
			log.Printf("CAN'T WRITE IN THE FILE %s", flake8File)
		}
	}
}

// GitHook set flake8 hook to git before commit
func GitHook(directory string) {
	command := fmt.Sprintf("cd %s && flake8 --install-hook git && git config --bool flake8.strict true", directory)
	output := packages.RunVenvCommand(command, directory)
	log.Printf("ADD HOOK FLAKE8: %s", output)
}
