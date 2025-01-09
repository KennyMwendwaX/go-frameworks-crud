package config

import (
	"log"

	"github.com/KennyMwendwaX/go-frameworks-crud/internal/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

type APIConfig struct {
	DB   *database.Queries
	pool *pgxpool.Pool
}

func ApiCfg() *APIConfig {
	pool, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	return &APIConfig{
		DB:   database.New(pool),
		pool: pool,
	}
}

func (cfg *APIConfig) Close() {
	if cfg.pool != nil {
		cfg.pool.Close()
	}
}
