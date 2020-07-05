package packages

import (
	"fmt"
	"log"
	"os"
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

	p.writeRequirements(p.installLibs(libs), "requirements.txt")
	p.writeRequirementsDev(p.installLibs(libsDev), "requirements-dev.txt")
}

func (p *Pip) installLibs(libs []string) []string {
	activate := path.Join(p.EnvDir, ".venv", "bin", "activate")
	log.Printf("ACTIVATE: %#v", activate)
	source := fmt.Sprintf("source %s && %s install %s && deactivate", activate, p.Name, strings.Join(libs, " "))
	log.Printf("SOURCE: %#v", source)
	cmd := exec.Command("bash", "-c", source)
	output, err := cmd.Output()
	utils.FatalError(err)

	return getLibsVersions(output)
}

func (p *Pip) writeRequirements(libs []string, fileName string) {
	f, err := os.OpenFile(path.Join(p.EnvDir, fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	utils.FatalError(err)
	defer f.Close()

	for _, lib := range libs {
		if _, err := f.WriteString(lib + "\n"); err != nil {
			log.Printf("ERROR WRITE LINE (%s): %s", fileName, lib)
		}
	}
}

func (p *Pip) writeRequirementsDev(libs []string, fileName string) {
	if _, err := os.Stat(path.Join(p.EnvDir, fileName)); os.IsNotExist(err) {
		f, err := os.OpenFile(path.Join(p.EnvDir, fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
		utils.FatalError(err)
		defer f.Close()

		firstLine := "-r requirements.txt"
		if _, err := f.WriteString(firstLine + "\n\n"); err != nil {
			log.Printf("ERROR WRITE LINE (%s): %s", fileName, firstLine)
		}
	}

	p.writeRequirements(libs, fileName)
}

// getLibsVersions get line of 'Succesfully installed' remove this term and return libs pinned with their versions
func getLibsVersions(output []byte) []string {
	lines := strings.Split(string(output), "\n")
	success := "Successfully installed "
	lastLine := lines[len(lines)-2]

	if !strings.HasPrefix(lastLine, success) {
		return []string{}
	}

	lastLine = strings.Replace(lastLine, success, "", 1)
	installedLibs := strings.Split(lastLine, " ")

	var libsPinned []string
	for _, lib := range installedLibs {
		libsPinned = append(libsPinned, pinLibVersion(lib))
	}

	return libsPinned
}

// pinLibVersion replace '-' to '==' and return pinned version
func pinLibVersion(lib string) string {
	log.Printf("LIB: %#v", lib)
	index := strings.LastIndexByte(lib, '-')
	return lib[:index] + "==" + lib[index+1:]
}
