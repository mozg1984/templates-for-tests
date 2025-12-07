package main

import (
	"context"

	"go-clickhouse/internal/config/clickhouse"
	clickhouseDB "go-clickhouse/internal/db/clickhouse"
	messagesRepo "go-clickhouse/internal/storage/messages"
	"go-clickhouse/internal/usecase"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := clickhouse.NewConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	db, err := clickhouseDB.New(cfg)
	if err != nil {
		logrus.Fatal(err)
	}

	defer func() {
		if err = db.Close(); err != nil {
			logrus.Error(err)
		}
	}()

	ctx := context.Background()

	if err = usecase.New(messagesRepo.New(db)).TestAddingMessages(ctx, 5000); err != nil {
		logrus.Fatal(err)
	}
}
