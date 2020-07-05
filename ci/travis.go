package ci

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
	travisTemplate = `language: python
python: %s
intall:
	- %s
script:
	- flake8
`
)

// CreateTravis create file '.travis.yml'
func CreateTravis(pack packages.Package) {
	travisFile := path.Join(pack.GetEnvDir(), ".travis.yml")

	travisContent := fmt.Sprintf(travisTemplate, packages.PythonVersion(), pack.GetCommand())
	log.Print(travisFile)
	log.Print(travisContent)
	if _, err := os.Stat(travisFile); os.IsNotExist(err) {
		if err := ioutil.WriteFile(travisFile, []byte(travisContent), 0664); err != nil {
			log.Printf("CAN'T WRITE IN THE FILE %s: %s", travisFile, err.Error())
		}
	}
}
