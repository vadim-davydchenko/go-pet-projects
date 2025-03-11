package database

import (
	"context"
	"go-pet-projects/fiber/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

func CreateDbPool(config *config.DatabaseConfig, logger *zerolog.Logger) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), config.Url)
	if err != nil {
		logger.Error().Msg("Unable to connect to database")
		panic(err)
	}
	logger.Info().Msg("Successfully connected to database")
	return dbpool
}
