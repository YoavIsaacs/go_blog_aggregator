package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DatabaseURL     string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigDirectory() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory")
	}

	return path + "/.gatorconfig.json", nil
}

func Read() (Config, error) {
	homeDir, err := getConfigDirectory()
	if err != nil {
		return Config{}, fmt.Errorf("error getting home directory")
	}

	byteValue, err := os.ReadFile(homeDir)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file")
	}

	var config Config

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error unmarshlling config file")
	}

	return config, nil
}

func (c *Config) SetUser(name string) error {
	homeDir, err := getConfigDirectory()
	if err != nil {
		return fmt.Errorf("error getting home directory")
	}

	c.CurrentUserName = name

	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling config")
	}

	err = os.WriteFile(homeDir, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing config to file")
	}

	return nil
}
