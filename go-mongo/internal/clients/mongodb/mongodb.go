package mongodb

import (
	"context"
	"fmt"
	"time"

	"go-mongo/internal/utils"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func NewClient(uri string, rp *readpref.ReadPref) (*mongo.Client, error) {
	bsonRegistry := bson.NewRegistry()
	bsonRegistry.RegisterTypeEncoder(utils.UUID, bson.ValueEncoderFunc(utils.EncodeUUIDValue))
	bsonRegistry.RegisterTypeDecoder(utils.UUID, bson.ValueDecoderFunc(utils.DecodeUUIDValue))

	client, err := mongo.Connect(options.
		Client().
		ApplyURI(uri).
		SetRegistry(bsonRegistry).
		SetReadPreference(rp),
	)
	if err != nil {
		return nil, fmt.Errorf("client creation error: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.Ping(ctx, rp); err != nil {
		return nil, fmt.Errorf("server is not responding: %w", err)
	}

	return client, nil
}
