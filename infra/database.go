package infra

import (
	"context"
	pgx "github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

var ConnectionPool *pgx.Pool

func InitDatabase(connectionString string, log *zap.Logger) error {
	pool, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		return err
	} else {
		ConnectionPool = pool
		return nil
	}
}

func CloseDatabase() {
	ConnectionPool.Close()
}

func GetDatabase() *pgx.Pool {
	return ConnectionPool
}
