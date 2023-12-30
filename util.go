package initenvs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from the specified files.
// If no paths are provided, add a nil, it loads from the default .env file.
// initenvs.LoadEnv()

func includesNil(paths []interface{}) bool {
	for _, path := range paths {
		if path == nil {
			return true
		}
	}
	return false
}

func LoadEnv(paths ...*string) {
	var (
		err     error
		my_path []any
	)

	for _, path := range paths {
		my_path = append(my_path, path)
	}

	if includesNil(my_path) || len(paths) == 0 {
		err = godotenv.Load()
	} else {
		for _, path := range paths {
			err = godotenv.Load(*path)
			if err != nil {
				log.Fatal(fmt.Sprintf("Error loading %s file", *path), err)
			}
		}
	}

	if err != nil {
		log.Fatal("Error loading .env file.", err)
	}
}

// GetEnv retrieves the value of the specified environment variable.
// If the variable is not set in your env file, it returns the default value.
// value := initenvs.GetEnv("KEY", "default")
func GetEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
