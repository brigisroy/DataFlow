package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"sample.code/dataflow/src/adapters/handler"
	"sample.code/dataflow/src/adapters/repository/postgres"
	"sample.code/dataflow/src/core/config"
	"sample.code/dataflow/src/core/service"
)

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	dbPool := postgres.CreateConnection()

	syncDataRepo := postgres.NewSyncDataRepository(dbPool)
	revenueRepo := postgres.NewRevenueRepository(dbPool)

	syncDataService := service.NewSyncDataService(config.App.DataFilePath, syncDataRepo)
	revenueService := service.NewRevenueService(revenueRepo)

	revenueHnadler := handler.NewRevenueHandler(revenueService)

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
		//After 30 minutes, trigger the sync data
		fmt.Println("30 minutes elapsed. Syncing data...")
		syncDataService.SyncData(ctx) // Trigger your sync process here
	}
	_ = syncDataService
	e := echo.New()

	// Register the route
	e.POST("/revenue", revenueHnadler.FetchRevenue)

	// Start the Echo server
	fmt.Println(fmt.Sprintf("Server is running on http://localhost:%s", config.App.Port))
	log.Fatal(e.Start(":" + config.App.Port))
}
