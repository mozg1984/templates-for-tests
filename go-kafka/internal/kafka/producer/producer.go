package producer

import (
	"errors"
	"fmt"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	flushTimeout = 5000 // ms
)

var ErrUnknownMessage = errors.New("unknown message")

type Encoder interface {
	Encode() ([]byte, error)
}

type SingleCallback func() error

type Producer struct {
	producer *kafka.Producer
}

func New(addresses []string) (*Producer, error) {
	cfg := &kafka.ConfigMap{
		"bootstrap.servers": strings.Join(addresses, ","),
	}

	producer, err := kafka.NewProducer(cfg)
	if err != nil {
		return nil, fmt.Errorf("error creating kafka producer: %w", err)
	}

	return &Producer{producer: producer}, nil
}

func (p *Producer) Produce(topic, key string, message Encoder) error {
	data, err := message.Encode()
	if err != nil {
		return fmt.Errorf("error encoding message: %w", err)
	}

	kafkaMsg := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: data,
		Key:   []byte(key), // murmur2 hash function
		//Headers: nil,
		//Opaque:  nil,
		//Timestamp:     time.Time{},
		//TimestampType: 0,
	}

	deliverChan := make(chan kafka.Event)

	if err = p.producer.Produce(&kafkaMsg, deliverChan); err != nil {
		return fmt.Errorf("error sending message: %w", err)
	}

	e := <-deliverChan
	switch ev := e.(type) {
	case kafka.Error:
		return ev
	case *kafka.Message:
		return nil
	default:
		return ErrUnknownMessage
	}
}

func (p *Producer) Close() {
	p.producer.Flush(flushTimeout)
	p.producer.Close()
}
