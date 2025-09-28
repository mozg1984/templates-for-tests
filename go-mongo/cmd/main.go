package main

import (
	"context"
	"fmt"
	"time"

	"go-mongo/internal/config"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go-mongo/internal/clients/mongodb"
	"go-mongo/internal/storage/users"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func main() {
	mongoDBConfig, err := config.MongoDBConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	client, err := mongodb.NewClient(mongoDBConfig.URI, readpref.SecondaryPreferred())
	if err != nil {
		logrus.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	repo := users.NewRepo(client)

	messages, err := repo.GetMessages(ctx, uuid.New())
	fmt.Println("messages: ", messages)
	fmt.Println("error: ", err)

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			logrus.Fatal(err)
		}
	}()
}
