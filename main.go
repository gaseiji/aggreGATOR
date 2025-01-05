package main

import (
	"aggregator/internal/command"
	"aggregator/internal/config"
	"aggregator/internal/state"
	"log"
	"os"
)

func main() {
	var st state.State

	configFile, err := config.ReadConfigFile()
	if err != nil {
		log.Fatalf("Error reading json config file %v", err)
	}

	st.Cfg = &configFile

	var cmdsStruct command.Commands

	cmdsStruct.Handlers = make(map[string]func(*state.State, command.Command) error)

	cmdsStruct.Register("login", command.HandlerLogin)

	if len(os.Args) < 2 {
		log.Fatalf("not enough arguments to execute. %v", err)
	}

	var cmd command.Command

	cmd.Name = os.Args[1]

	for i, v := range os.Args {
		if i > 1 {
			cmd.Args = append(cmd.Args, v)
		}
	}

	err = cmdsStruct.Run(&st, cmd)
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}

	//fmt.Printf("db:%s current user:%s\n", configStruct.DbUrl, configStruct.CurrentUserName)

}
