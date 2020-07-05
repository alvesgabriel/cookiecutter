package packages

import (
	"fmt"
	"log"
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

	p.installLibs()
}

func (p *Pip) installLibs() {
	activate := path.Join(p.EnvDir, ".venv", "bin", "activate")
	log.Printf("ACTIVATE: %#v", activate)
	source := fmt.Sprintf("source %s && %s install %s && deactivate", activate, p.Name, strings.Join(libs, " "))
	log.Printf("SOURCE: %#v", source)
	cmd := exec.Command("bash", "-c", source)
	output, err := cmd.Output()
	utils.FatalError(err)
	log.Printf("%s", output)
}
