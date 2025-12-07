package models

import "time"

type Message struct {
	UserID uint64    `json:"user_id"`
	Text   string    `json:"text"`
	Date   time.Time `json:"date"`
}

type MessagesCount struct {
	UserID uint64 `json:"user_id"`
	Count  uint64 `json:"count"`
}
