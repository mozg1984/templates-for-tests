package consumer

import (
	"context"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

const (
	sessionTimeout = 7000 // ms
	noTimeout      = -1   // blocking waiting
)

type Handler interface {
	Handle(ctx context.Context, message *kafka.Message, consumerId int) error
}

type Consumer struct {
	consumer *kafka.Consumer
	handler  Handler
	stop     bool
	id       int
}

func rebalanceCb(*kafka.Consumer, kafka.Event) error {
	// TODO
	return nil
}

func New(address []string, topic, groupID string, id int, handler Handler) (*Consumer, error) {
	cfg := kafka.ConfigMap{
		"bootstrap.servers":        strings.Join(address, ","),
		"group.id":                 groupID,
		"session.timeout.ms":       sessionTimeout,
		"enable.auto.offset.store": false, // default = true (It's responsible for storing offsets in consumer's memory)
		"enable.auto.commit":       true,
		"auto.commit.interval.ms":  sessionTimeout,
		"auto.offset.reset":        "earliest",
	}

	consumer, err := kafka.NewConsumer(&cfg)
	if err != nil {
		return nil, err
	}

	if err = consumer.Subscribe(topic, rebalanceCb); err != nil {
		return nil, err
	}

	return &Consumer{
		consumer: consumer,
		handler:  handler,
		id:       id,
	}, nil
}

func (c *Consumer) Start(ctx context.Context) {
	for {
		if c.stop {
			break
		}

		kafkaMsg, err := c.consumer.ReadMessage(noTimeout)
		if err != nil {
			logrus.Error(err)
		}
		if kafkaMsg == nil {
			continue
		}

		if err = c.handler.Handle(ctx, kafkaMsg, c.id); err != nil {
			logrus.Error(err)
			continue
		}

		if _, err = c.consumer.StoreMessage(kafkaMsg); err != nil {
			logrus.Error(err)
			continue
		}
	}
}

func (c *Consumer) Close() error {
	c.stop = true

	// since we store offsets manually we should ensure before closing
	// that all stored offsets are committed in kafka
	if _, err := c.consumer.Commit(); err != nil {
		return err
	}

	return c.consumer.Close()
}
