package packages

import (
	"os/exec"
	"path"
	"strings"

	// Local
	"github.com/alvesgabriel/cookiecutter/utils"
)

// Pip mangager package
type Pip struct {
	Name   string
	EnvDir string
}

// Create inicialize manager package using pip
func (p Pip) Create() {
	python := "python3"
	direcotories := strings.Split(p.EnvDir, "/")
	project := direcotories[len(direcotories)-1]
	cmd := exec.Command(python, "-m", "venv", path.Join(p.EnvDir, ".venv"), "--prompt", project)
	err := cmd.Run()
	utils.FatalError(err)
}
