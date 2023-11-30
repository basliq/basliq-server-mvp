package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

type Config struct {
	Username string
	Password string
	Port     int
	Host     string
	DBName   string
}

type DB struct {
	config Config
	db     *pgxpool.Pool
}

func New(config Config) *DB {
	db, err := pgxpool.New(context.Background(), fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", config.Username, config.Password, config.Host, config.Port, config.DBName))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	return &DB{
		config: config,
		db:     db,
	}
}
