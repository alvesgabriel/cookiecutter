package main

import (
	"log"
	"os"

	// Third-Party
	"github.com/galdor/go-cmdline"

	// Local
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

	directory = pathDir
	pack = "pip"

	cmd := cmdline.New()
	cmd.AddOption("d", "directory", "dir", "Project directory")
	cmd.SetOptionDefault("d", "current directory")
	cmd.AddOption("p", "manager_packages", "pack", "Manager package to install dependences")
	cmd.SetOptionDefault("p", pack)

	cmd.Parse(os.Args)

	if cmd.IsOptionSet("d") {
		directory = cmd.OptionValue("d")
	}
	if cmd.IsOptionSet("p") {
		pack = cmd.OptionValue("p")
	}

	var packs = map[string]packages.Package{
		"pip": packages.Pip{"pip", directory},
	}

	managerPackage = packs[pack]

	createDir(directory)
	packages.CreateVenv(managerPackage)
}

func createDir(directory string) {
	log.Printf("CREATING DIR: %#v", directory)
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		os.Mkdir(directory, os.FileMode(0775))
	}
	repository.Init(directory)
}
