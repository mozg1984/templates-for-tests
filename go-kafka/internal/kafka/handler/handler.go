package handler

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx context.Context, msg *kafka.Message, consumerID int) error {
	logrus.WithContext(ctx).Infof("Consumer #%d: message from kafka with offset %d '%s' on partition %d",
		consumerID, msg.TopicPartition.Offset, msg.Value, msg.TopicPartition.Partition)
	return nil
}
