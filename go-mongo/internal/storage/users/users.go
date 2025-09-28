package users

import (
	"context"

	"go-mongo/internal/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repo struct {
	usersDB *mongo.Database
}

const (
	usersCollection    = "users"
	messagesCollection = "users"
)

func NewRepo(client *mongo.Client) *Repo {
	return &Repo{
		usersDB: client.Database(usersCollection),
	}
}

func (r *Repo) GetMessages(ctx context.Context, userID uuid.UUID) (messages []models.Message, err error) {
	cursor, err := r.usersDB.Collection(messagesCollection).Find(ctx, bson.M{"_id": bson.M{"$eq": userID}})
	if err != nil {
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var m models.Message

		err = cursor.Decode(&m)
		if err != nil {
			return
		}

		messages = append(messages, m)
	}

	return messages, cursor.Err()
}
