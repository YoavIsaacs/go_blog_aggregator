package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	DatabaseURL string `json:"db_url"`
}

func Read() Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error getting home directory...")
	}

	byteValue, err := os.ReadFile(homeDir + ".gatorconfig")
	if err != nil {
		log.Fatal("Error reading config file...")
	}

	var config Config

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatal("Error unmarshlling config file...")
	}

	fmt.Println(config.DatabaseURL)
	return config
}
