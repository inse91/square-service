package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"square-service/internal/config"
)

func NewPostgresClient(ctx context.Context, cfg config.PostgresDB) (*pgx.Conn, error) {

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DataBase)

	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		return nil, errors.Wrap(err,
			fmt.Sprintf(
				"failed to set connection to postgres with connection string: %s",
				connectionString,
			))
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to ping postgres")
	}

	return conn, nil

}
