package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var App AppConfig

type AppConfig struct {
	DatabaseURL  string
	DataFilePath string
	Port         string
}

func init() {
	loadConfig()

	fmt.Println("All configuration variables are set.")
}

// loadConfig loads the environment variables.
// It exits the program if any variable is not set.
func loadConfig() {
	App = AppConfig{}
	var missingVars []string

	configMap := map[string]*string{
		"DATABASE_URL":   &App.DatabaseURL,
		"DATA_FILE_PATH": &App.DataFilePath,
		"PORT":           &App.Port,
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		os.Exit(1)
	}

	for key, field := range configMap {
		value := os.Getenv(key)
		if value == "" {
			missingVars = append(missingVars, key)
		} else {
			*field = value
		}
	}

	if len(missingVars) != 0 {
		_, _ = fmt.Printf("The following configuration variables are not set: %v\n", missingVars)
		os.Exit(1)
	}
}
