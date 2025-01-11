package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"sample.code/dataflow/src/core/config"
)

func CreateConnection() *pgxpool.Pool {
	connStr := config.App.DatabaseURL
	var err error
	dbPool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
	}
	return dbPool
}
