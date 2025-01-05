package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"
const configFilePath = "/home/gabrielseji/projects/aggreGATOR"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	return configFilePath + "/" + configFileName, nil
}

func ReadConfigFile() (Config, error) {
	filepath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	jsondata, err := os.ReadFile(filepath)
	if err != nil {
		return Config{}, err
	}
	configStruct := Config{}
	err = json.Unmarshal(jsondata, &configStruct)
	if err != nil {
		return Config{}, err
	}
	return configStruct, nil
}

func (c *Config) SetUser(user string) error {
	configStruct, err := ReadConfigFile()
	if err != nil {
		return err
	}
	configStruct.CurrentUserName = user
	write(configStruct)
	fmt.Println("User set sucessfully")
	return nil
}

func write(cfg Config) error {
	jsondata, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	filepath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath, jsondata, 0644)
	if err != nil {
		return err
	}
	return nil
}
