package state

import (
	"aggregator/internal/config"
	"aggregator/internal/database"
)

type State struct {
	Cfg *config.Config
	Db  *database.Queries
}
