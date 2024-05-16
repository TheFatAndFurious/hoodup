package utils

import (
	"log"
	"os"
	"path/filepath"
)




func GetAbsolutePath(relativePath string) string {
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	return filepath.Join(rootPath, relativePath)
}