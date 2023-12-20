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
	pool   *pgxpool.Pool
}

func New(config Config) *DB {
	pool, err := pgxpool.New(context.Background(), fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", config.Username, config.Password, config.Host, config.Port, config.DBName))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	// TODO - where to put the close?
	// defer pool.Close(context.Background())

	fmt.Println("connected to database:", &pool)
	return &DB{
		config: config,
		pool:   pool,
	}
}

func (db DB) Ping() {
	err := db.pool.Ping(context.Background())
	if err != nil {
		fmt.Println("there is no connection:", err)
	}
}
