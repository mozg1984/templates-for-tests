package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go-kafka/internal/config"
	kafkaConsumer "go-kafka/internal/kafka/consumer"
	kafkaHandler "go-kafka/internal/kafka/handler"

	"github.com/sirupsen/logrus"
)

const (
	topic   = "test-items"
	groupID = "test-items-consumer-group"
)

func main() {
	ctx := context.Background()

	kafkaConfig, err := config.KafkaConfig()
	if err != nil {
		logrus.Fatal("Kafka config getting error: ", err)
	}

	handler := kafkaHandler.New()
	consumer1, err := kafkaConsumer.New(kafkaConfig.Brokers, topic, groupID, 1, handler)
	if err != nil {
		logrus.Fatal(err)
	}

	consumer2, err := kafkaConsumer.New(kafkaConfig.Brokers, topic, groupID, 2, handler)
	if err != nil {
		logrus.Fatal(err)
	}

	consumer3, err := kafkaConsumer.New(kafkaConfig.Brokers, topic, groupID, 3, handler)
	if err != nil {
		logrus.Fatal(err)
	}

	consumer4, err := kafkaConsumer.New(kafkaConfig.Brokers, topic, groupID, 4, handler)
	if err != nil {
		logrus.Fatal(err)
	}

	go func() {
		consumer1.Start(ctx)
	}()

	go func() {
		consumer2.Start(ctx)
	}()

	go func() {
		consumer3.Start(ctx)
	}()

	go func() {
		consumer4.Start(ctx)
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	if err = consumer1.Close(); err != nil {
		logrus.Fatal(err)
	}

	if err = consumer2.Close(); err != nil {
		logrus.Fatal(err)
	}

	if err = consumer3.Close(); err != nil {
		logrus.Fatal(err)
	}

	if err = consumer4.Close(); err != nil {
		logrus.Fatal(err)
	}
}
