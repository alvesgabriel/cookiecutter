package lint

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	// Local
	"github.com/alvesgabriel/cookiecutter/packages"
)

const (
	flake8Content = `[flake8]
max-line-length=120
exclude=.venv`
)

// CreateFlake8 create file '.flake8'
func CreateFlake8(directory string) {
	flake8File := path.Join(directory, ".flake8")

	if _, err := os.Stat(flake8File); os.IsNotExist(err) {
		if err := ioutil.WriteFile(flake8File, []byte(flake8Content), 0664); err != nil {
			log.Printf("CAN'T WRITE IN THE FILE %s: %s", flake8File, err.Error())
		}
	}
}

// GitHook set flake8 hook to git before commit
func GitHook(directory string) {
	command := fmt.Sprintf("cd %s && flake8 --install-hook git && git config --bool flake8.strict true", directory)
	output := packages.RunVenvCommand(command, directory)
	log.Printf("ADD HOOK FLAKE8: %s", output)
}
