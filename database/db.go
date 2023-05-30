package database

import (
	"context"
	"fmt"
	"os"
	"project/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Init_db() {
	pool, err := pgxpool.New(context.Background(), config.ENV.DB_URI)

	if err != nil {
		fmt.Println("Unable to create connection pool: ", err.Error())
		os.Exit(1)
	}

	DB = pool
}
