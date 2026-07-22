package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)


const configFilename = ".gatorconfig.json"

type Config struct {
	DbUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	return write(*c)
}

func getConfigFilepath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, configFilename), nil
}

func Read() (Config, error) {
	filePath, err := getConfigFilepath()
	cfg := Config{}
	if err != nil {
		return cfg, err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return cfg, err
	}
	defer file.Close()
	
	err = json.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}

func write(cfg Config) error {
	filePath, err := getConfigFilepath()
	if err != nil {
		return err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}



