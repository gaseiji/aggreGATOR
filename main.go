package main

import (
	"aggregator/internal/config"
	"fmt"
	"log"
)

func main() {
	configStruct, err := config.ReadConfigFile()

	config.SetUser("Gabriel")

	configStruct, err = config.ReadConfigFile()

	if err != nil {
		log.Fatalf("Error reading json config file %v", err)
	}

	fmt.Printf("db:%s current user:%s\n", configStruct.DbUrl, configStruct.CurrentUserName)
}
