package utils

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}




func InitEnv() {
	// Debugging: Print current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	log.Printf("Current working directory: %s", cwd)

	// Debugging: List files in the current directory
	files, err := ioutil.ReadDir(cwd)
	if err != nil {
		log.Fatalf("Error reading current directory: %v", err)
	}
	for _, file := range files {
		log.Printf("File: %s", file.Name())
	}

	// Load .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}



