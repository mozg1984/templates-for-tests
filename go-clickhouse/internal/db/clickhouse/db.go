package clickhouse

import (
	"strings"
	"time"

	ch "go-clickhouse/internal/config/clickhouse"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type DB struct {
	driver.Conn
}

func New(cfg ch.Config) (*DB, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: strings.Split(cfg.Endpoints, ","),
		Auth: clickhouse.Auth{
			Database: cfg.DBName,
			Username: cfg.User,
			Password: cfg.Password,
		},
		//Debug: true,
		//Debugf: func(format string, v ...interface{}) {
		//	logrus.Debugf(format, v...)
		//},
		//DialContext: func(ctx context.Context, addr string) (net.Conn, error) {
		//	var d net.Dialer
		//	return d.DialContext(ctx, "tcp", addr)
		//},
		Settings: clickhouse.Settings{
			"max_execution_time": 5,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		DialTimeout:      time.Duration(10) * time.Second,
		MaxOpenConns:     20,
		MaxIdleConns:     20,
		ConnMaxLifetime:  time.Duration(5) * time.Minute,
		ConnOpenStrategy: clickhouse.ConnOpenInOrder,
		BlockBufferSize:  10,
	})
	if err != nil {
		return nil, err
	}

	return &DB{conn}, nil
}
