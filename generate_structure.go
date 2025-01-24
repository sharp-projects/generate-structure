package main

import "fmt"

// list of directories
var directories = []string{
	"cmd",
	"cmd/app",
	"configs",
	"internal/handlers",
	"internal/services",
	"internal/repositories",
	"internal/models",
	"pkg/logger",
	"pkg/config",
	"pkg/errors",
	"migrations",
	"tests/integration",
	"tests/unit",
	"tests/mocks",
	"docs",
}

// list of files
var files = []string{
	"README.md",
	"LICENSE",
	".gitignore",
	"Makefile",
	"Dockerfile",
	"docker-compose.yml",
	".env",
}

// function to create the directories
func createDirectories() {
	fmt.Println("Creating directories")
}

// function to create the files
func createFiles() {
	fmt.Println("Creating files")
}

// function to validate that all directories and files were created
func validateStructure() {
	fmt.Println("Validating structure")
}

func main() {
	createDirectories()
	createFiles()
	validateStructure()
	// create directories
	// create files
	// validate the structure
}
