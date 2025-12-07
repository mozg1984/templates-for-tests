package messages

import (
	"context"
	"fmt"
	"time"

	"go-clickhouse/internal/db/clickhouse"
	"go-clickhouse/internal/models"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/rs/zerolog/log"
)

type Repo struct {
	db *clickhouse.DB
}

func New(db *clickhouse.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Add(messages []models.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := "INSERT INTO ch.messages (user_id, text, date) VALUES (?, ?, ?)"

	batch, err := r.db.PrepareBatch(ctx, query)
	if err != nil {
		return err
	}

	for _, m := range messages {
		err = batch.Append(
			m.UserID,
			m.Text,
			m.Date,
		)
		if err != nil {
			return err
		}
	}

	return batch.Send()
}

func (r *Repo) GetCounts(ctx context.Context, limit int) (messagesCount []models.MessagesCount, err error) {
	query := "SELECT user_id, count() FROM ch.messages GROUP BY user_id LIMIT ?"

	rows, err := r.db.Query(ctx, query, limit)
	if err != nil {
		return messagesCount, fmt.Errorf("error occured while getting messages count: %w", err)
	}
	defer func(rows driver.Rows) {
		err := rows.Close()
		if err != nil {
			log.Error().Err(err).Msg("close rows error")
		}
	}(rows)

	for rows.Next() {
		var msgCount models.MessagesCount

		if err = rows.Scan(&msgCount.UserID, &msgCount.Count); err != nil {
			err = fmt.Errorf("scan messages count error: %w", err)
			return
		}

		messagesCount = append(messagesCount, msgCount)
	}

	return
}
