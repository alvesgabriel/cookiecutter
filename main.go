package main

import (
	"log"
	"os"

	// Third-Party
	"github.com/galdor/go-cmdline"

	// Local
	"github.com/alvesgabriel/cookiecutter/ci"
	"github.com/alvesgabriel/cookiecutter/lint"
	"github.com/alvesgabriel/cookiecutter/packages"
	"github.com/alvesgabriel/cookiecutter/repository"
	"github.com/alvesgabriel/cookiecutter/utils"
)

func main() {
	var (
		directory      string
		pack           string
		managerPackage packages.Package
	)

	pathDir, err := os.Getwd()
	utils.FatalError(err)

	// Default options CLI
	directory = pathDir
	pack = "pip"
	ciService := "travis"

	cmd := cmdline.New()
	cmd.AddOption("d", "directory", "dir", "Project directory")
	cmd.SetOptionDefault("d", "current directory")
	cmd.AddOption("p", "manager-packages", "pack", "Manager package to install dependences")
	cmd.SetOptionDefault("p", pack)
	cmd.AddOption("c", "continuous-integration", "continuous integration", "Continuous integration service")
	cmd.SetOptionDefault("c", ciService)

	cmd.Parse(os.Args)

	if cmd.IsOptionSet("d") {
		directory = cmd.OptionValue("d")
	}
	if cmd.IsOptionSet("p") {
		pack = cmd.OptionValue("p")
	}
	if cmd.IsOptionSet("c") {
		ciService = cmd.OptionValue("c")
	}

	var packs = map[string]packages.Package{
		"pip": packages.NewPip(directory),
	}

	managerPackage = packs[pack]

	createDir(directory)
	packages.CreateVenv(managerPackage)
	lint.CreateFlake8(directory)
	lint.GitHook(directory)
	ci.CreateTravis(managerPackage)
}

func createDir(directory string) {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		log.Printf("CREATING DIR: %#v", directory)
		os.Mkdir(directory, os.FileMode(0775))
	}
	repository.Init(directory)
}
