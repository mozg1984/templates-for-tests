package main

import (
	"fmt"

	"go-kafka/internal/config"
	kafkaProducer "go-kafka/internal/kafka/producer"
	"go-kafka/internal/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	topic = "test-items"
)

func main() {
	kafkaConfig, err := config.KafkaConfig()
	if err != nil {
		logrus.Fatal("Kafka config getting error: ", err)
	}

	producer, err := kafkaProducer.New(kafkaConfig.Brokers)
	if err != nil {
		logrus.Fatal("Kafka producer creation error: ", err)
	}

	for _, item := range testItems(100) {
		if err = producer.Produce(topic, item.Version, item); err != nil {
			logrus.Error("Kafka producer error: ", err)
		}
	}

	for _, item := range testItems(100) {
		if err = producer.Produce(topic, item.Version, item); err != nil {
			logrus.Error("Kafka producer error: ", err)
		}
	}

	for _, item := range testItems(100) {
		if err = producer.Produce(topic, item.Version, item); err != nil {
			logrus.Error("Kafka producer error: ", err)
		}
	}

	producer.Close()
}

func testItems(count int) []*models.Item {
	version := uuid.NewString()

	items := make([]*models.Item, count)
	for i := 0; i < count; i++ {
		items[i] = &models.Item{
			ID:          int64(i + 1),
			Name:        fmt.Sprintf("Item #%d", i+1),
			Description: fmt.Sprintf("This is a test item #%d", i+1),
			Version:     version,
		}
	}

	return items
}
