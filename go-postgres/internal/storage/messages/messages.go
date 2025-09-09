package messages

import (
	"context"
	"fmt"

	"go-postgres/internal/db"
	"go-postgres/internal/models"

	"github.com/jackc/pgx/v5"
)

type Repo struct {
	db *db.DB
}

func New(db *db.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Add(ctx context.Context, msg models.Message) (msgID int64, err error) {
	query := `SELECT add_message(p_value := $1, p_author := $2)`

	if err = r.db.Pool().QueryRow(ctx, query, msg.Value, msg.Author).Scan(&msgID); err != nil {
		return 0, fmt.Errorf("failed while adding message: %w", err)
	}

	return
}

func (r *Repo) GetAll(ctx context.Context, limit int) (messages []models.Message, err error) {
	query := `SELECT id, value, author FROM messages ORDER BY id LIMIT $1`

	var rows pgx.Rows
	if rows, err = r.db.Pool().Query(ctx, query, limit); err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var row models.Message

		if err = rows.Scan(
			&row.ID,
			&row.Value,
			&row.Author,
		); err != nil {
			return
		}

		messages = append(messages, row)
	}

	return messages, rows.Err()
}

func (r *Repo) DeleteAll(ctx context.Context) (err error) {
	_, err = r.db.Pool().Exec(ctx, "DELETE FROM messages")

	return
}
