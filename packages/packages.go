package packages

import (
	"log"
	"os/exec"
	"strings"

	"github.com/alvesgabriel/cookiecutter/utils"
)

// Package interface to manager package
type Package interface {
	Create()
	GetEnvDir() string
	GetCommand() string
}

var (
	libs    = []string{"requests"}
	libsDev = []string{"flake8", "black"}
)

// CreateVenv creates virtualenv
func CreateVenv(pack Package) {
	log.Printf("CREATING VIRTUALENV")
	pack.Create()
}

// PythonVersion get version of Python
func PythonVersion() string {
	cmd := exec.Command("python", "--version")
	output, err := cmd.Output()
	utils.FatalError(err)
	return strings.Split(string(output), " ")[1]
}
