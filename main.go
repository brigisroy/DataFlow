package main

import (
	"log"

	"sample.code/dataflow/src/adapters/repository/postgres"
	"sample.code/dataflow/src/core/config"
	"sample.code/dataflow/src/core/service"
)

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	dbPool := postgres.CreateConnection()

	syncDataRepo := postgres.NewSyncDataRepository(dbPool)
	syncDataService := service.NewSyncDataService(config.App.DataFilePath, syncDataRepo)

	syncDataService.SyncData()
}
