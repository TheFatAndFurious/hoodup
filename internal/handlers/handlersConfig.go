// internal/handlers/handlers.go
package handlers

import (
	"log"
	"os"

	"goserver.com/internal/utils"
)

var basePath string

// InitializeHandlers sets up necessary configurations for handlers
func InitializeHandlers() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}
	log.Printf("Current working directory: %s", cwd)
	relativePath := utils.GetEnv("TEMPLATES_PATH", "web/templates/")
	log.Printf("TEMPLATES_PATH: %s", relativePath)
	basePath = utils.GetAbsolutePath(relativePath)
	log.Printf("Template path set to: %s", basePath)
	files, err := os.ReadDir(basePath)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}
	for _, file := range files {
		log.Printf("Found file: %s", file.Name())
	}
}
