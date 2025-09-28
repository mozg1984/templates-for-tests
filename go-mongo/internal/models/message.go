package models

import "github.com/google/uuid"

type Message struct {
	ID      uuid.UUID `bson:"_id,omitempty"`
	Title   string    `bson:"title"`
	Content string    `bson:"content"`
}
