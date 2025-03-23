package main

import (
	"fmt"
	"github.com/DataStarETL/Orchestrator/internal/config"
	"github.com/DataStarETL/Orchestrator/internal/database"
	"github.com/DataStarETL/Orchestrator/internal/rest"
)

func main() {
	databaseConfig, serverConfig := config.LoadConfig()
	db, err := database.CreateDBConnection(databaseConfig)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := rest.CreateRestService(db)
	err = r.Start(fmt.Sprintf(":%s", serverConfig.Port))
	if err != nil {
		panic(err)
	}
}
