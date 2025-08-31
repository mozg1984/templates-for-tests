package models

import (
	"fmt"

	jsonIter "github.com/json-iterator/go"
)

type Item struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

func (i *Item) Encode() (data []byte, err error) {
	if data, err = jsonIter.Marshal(i); err != nil {
		err = fmt.Errorf("failed to encode item to JSON: %w", err)
		return
	}

	return
}
