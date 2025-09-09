package db

import (
	"context"
	"fmt"
	"time"

	"go-postgres/internal/config"

	pgxLogrus "github.com/jackc/pgx-logrus"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/sirupsen/logrus"
)

type DB struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, cfg config.DBConfig) (*DB, error) {
	pool, err := newPool(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &DB{pool: pool}, nil
}

func newPool(ctx context.Context, cfg config.DBConfig) (*pgxpool.Pool, error) {
	pgxConnConfig, err := pgxpool.ParseConfig(cfg.ConnString())
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	pgxConnConfig.MaxConns = cfg.MaxConnections
	pgxConnConfig.MaxConnLifetime = time.Second * time.Duration(cfg.MaxConnectionLifetimeSec)
	pgxConnConfig.MaxConnIdleTime = time.Second * time.Duration(cfg.MaxConnectionIdleTimeSec)

	if cfg.EnableLogging {
		l := logrus.New()
		l.SetLevel(logrus.DebugLevel)
		l.SetFormatter(new(logrus.JSONFormatter))

		pgxConnConfig.ConnConfig.Tracer = &tracelog.TraceLog{
			Logger:   pgxLogrus.NewLogger(l),
			LogLevel: tracelog.LogLevelDebug,
		}
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), pgxConnConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return pool, nil
}

func (c DB) Pool() *pgxpool.Pool {
	return c.pool
}

func (c DB) Close() {
	c.pool.Close()
}
