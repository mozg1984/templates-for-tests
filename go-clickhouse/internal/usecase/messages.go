package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go-clickhouse/internal/models"
)

type messagesRepo interface {
	Add([]models.Message) error
	GetCounts(context.Context, int) ([]models.MessagesCount, error)
}

type UseCase struct {
	repo messagesRepo
}

func New(repo messagesRepo) *UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) TestAddingMessages(ctx context.Context, msgCount int) error {
	if msgCount <= 0 {
		return errors.New("msg count must be greater than zero")
	}

	userIDs := []uint64{
		1234567891,
		1234567892,
		1234567893,
		1234567894,
		1234567895,
	}

	texts := []string{
		"Hi",
		"Hey",
		"Hello",
	}

	messages := make([]models.Message, 0, msgCount)
	for i := 0; i < msgCount; i++ {
		messages = append(messages, models.Message{
			UserID: userIDs[i%len(userIDs)],
			Text:   texts[i%len(texts)],
			Date:   time.Now(),
		})
	}

	if err := u.repo.Add(messages); err != nil {
		return fmt.Errorf("adding messages error: %w", err)
	}

	messagesCounts, err := u.repo.GetCounts(ctx, len(userIDs))
	if err != nil {
		return fmt.Errorf("getting counts error: %w", err)
	}

	if len(messagesCounts) != len(userIDs) {
		return fmt.Errorf("the size of messages counts (%d) is not equal to the size of user ids (%d)",
			len(messagesCounts), len(userIDs))
	}

	return nil
}
