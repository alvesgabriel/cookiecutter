package packages

import "log"

// Package interface to manager package
type Package interface {
	Create()
}

var (
	libs    = []string{"requests"}
	libsDev = []string{"flake8"}
)

// CreateVenv creates virtualenv
func CreateVenv(pack Package) {
	log.Printf("CREATING VIRTUALENV")
	pack.Create()
}
