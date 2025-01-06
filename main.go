package main

import (
	"aggregator/internal/command"
	"aggregator/internal/config"
	"aggregator/internal/database"
	"aggregator/internal/state"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	var st state.State

	configFile, err := config.ReadConfigFile()
	if err != nil {
		log.Fatalf("Error reading json config file %v", err)
	}
	st.Cfg = &configFile

	db, err := sql.Open("postgres", st.Cfg.DbUrl)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging the database:", err)
	}

	dbQueries := database.New(db)
	st.Db = dbQueries

	var cmdsStruct command.Commands
	cmdsStruct.Handlers = make(map[string]func(*state.State, command.Command) error)
	cmdsStruct.Register("login", command.HandlerLogin)
	cmdsStruct.Register("register", command.HandlerRegister)
	cmdsStruct.Register("reset", command.HandlerResetDb)
	cmdsStruct.Register("users", command.HandlerUsers)

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

}
