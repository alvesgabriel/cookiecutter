package packages

import "log"

// Package interface to manager package
type Package interface {
	Create()
}

// CreateVenv creates virtualenv
func CreateVenv(pack Package) {
	log.Printf("CREATING VIRTUALENV")
	pack.Create()
}
