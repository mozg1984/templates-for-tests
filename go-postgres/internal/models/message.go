package models

type Message struct {
	ID     int64   `json:"id"`
	Value  string  `json:"value"`
	Author *string `json:"author"`
}
