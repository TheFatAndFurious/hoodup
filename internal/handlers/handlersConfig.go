// internal/handlers/handlers.go
package handlers

import (
	"log"

	"goserver.com/internal/utils"
)

var basePath string

// InitializeHandlers sets up necessary configurations for handlers
func InitializeHandlers() {
	relativePath := utils.GetEnv("TEMPLATES_PATH", "web/templates/")
	basePath = utils.GetAbsolutePath(relativePath)
	log.Printf("Template path set to: %s", basePath)
}
