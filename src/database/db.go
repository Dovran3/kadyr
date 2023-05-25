package database

import (
	"admin/src/loadEnv"
	"os"
	"fmt"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)
var DB *pgxpool.Pool

func Init_db() {
	loadEnv.LoadEnvVariables()
	pool, err := pgxpool.New(context.Background(),os.Getenv("DB_URI"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	DB = pool
}
