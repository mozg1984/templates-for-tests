package main

import (
	"context"

	"go-postgres/internal/config"
	"go-postgres/internal/db"
	messagesRepo "go-postgres/internal/storage/messages"
	messagesUseCase "go-postgres/internal/usecase/messages"

	"github.com/sirupsen/logrus"
)

func main() {
	dbConfig, err := config.NewDBConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	ctx := context.Background()

	dbInstance, err := db.New(ctx, dbConfig)
	if err != nil {
		logrus.Fatal(err)
	}

	defer dbInstance.Close()

	repo := messagesRepo.New(dbInstance)
	useCase := messagesUseCase.New(repo)

	if err = useCase.TestAddingMessages(ctx); err != nil {
		logrus.Fatal(err)
	}
}
