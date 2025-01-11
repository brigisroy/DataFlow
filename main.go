package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"sample.code/dataflow/src/adapters/repository/postgres"
	"sample.code/dataflow/src/core/config"
	"sample.code/dataflow/src/core/service"
)

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	dbPool := postgres.CreateConnection()

	syncDataRepo := postgres.NewSyncDataRepository(dbPool)
	syncDataService := service.NewSyncDataService(config.App.DataFilePath, syncDataRepo)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ticketChannel := make(chan bool, 1)
	ticker := time.NewTicker(30 * time.Minute)
	defer ticker.Stop()
	go func() {
		// Wait for the ticker to expire and signal the syncData method
		<-ticker.C
		// Ticket for 30 minutes elapsed, trigger sync data
		ticketChannel <- true
	}()

	// Wait for the ticket signal and call SyncData
	select {
	case <-ticketChannel:
		// After 30 minutes, trigger the sync data
		fmt.Println("30 minutes elapsed. Syncing data...")
		syncDataService.SyncData(ctx) // Trigger your sync process here
	}
}
