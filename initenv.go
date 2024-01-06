package initenv

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from the specified files.
// If no paths are provided, it loads from the default .env file.
//
// It logs a fatal error if any file fails to load.
func LoadEnv(paths ...string) {
	var err error

	if len(paths) == 0 {
		err = godotenv.Load()
	} else {
		for _, path := range paths {
			err = godotenv.Load(path)
			if err != nil {
				log.Fatal(fmt.Sprintf("Error loading %s file", path), err)
			}
		}
	}

	if err != nil {
		log.Fatal("Error loading .env file.", err)
	}
}

// GetEnv retrieves the value of the specified environment variable.
// If the variable is not set in your env file, it returns the default value.
//
// Example:
//
//	value := initenv.GetEnv("KEY", "default")
func GetEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
