package config

import (
	"encoding/json"
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
