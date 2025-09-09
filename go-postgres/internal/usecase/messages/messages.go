package messages

import (
	"context"
	"fmt"
	"slices"

	"go-postgres/internal/models"
)

const testMessagesCount = 10

type messagesRepo interface {
	Add(ctx context.Context, msg models.Message) (msgID int64, err error)
	GetAll(ctx context.Context, limit int) (messages []models.Message, err error)
	DeleteAll(ctx context.Context) (err error)
}

type UseCase struct {
	repo messagesRepo
}

func New(repo messagesRepo) *UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) TestAddingMessages(ctx context.Context) (err error) {
	msgIDs1 := make([]int64, 0, testMessagesCount)

	// step 1. Reset all previous messages
	if err = u.repo.DeleteAll(ctx); err != nil {
		err = fmt.Errorf("failed while pre deleting all message: %w", err)
		return
	}

	// step 2. Add new messages
	for i := range testMessagesCount {
		author := fmt.Sprintf("tester-%d", i)

		msg := models.Message{
			Value:  fmt.Sprintf("message-%d", i),
			Author: &author,
		}

		var msgID int64
		if msgID, err = u.repo.Add(ctx, msg); err != nil {
			err = fmt.Errorf("failed while pre deleting all message: %w", err)
			return
		}

		msgIDs1 = append(msgIDs1, msgID)
	}

	// step 3. Get all recently added messages
	messages, err := u.repo.GetAll(ctx, 2*testMessagesCount)
	if err != nil {
		err = fmt.Errorf("failed while pre deleting all message: %w", err)
		return
	}

	// step 4. Reset all recently added messages
	if err = u.repo.DeleteAll(ctx); err != nil {
		err = fmt.Errorf("failed while post deleting all message: %w", err)
		return
	}

	// step 5. Prepare data
	msgIDs2 := make([]int64, 0, len(messages))
	for _, m := range messages {
		msgIDs2 = append(msgIDs2, m.ID)
	}

	fmt.Println("msg IDs:")
	fmt.Println(msgIDs1)
	fmt.Println(msgIDs2)

	// step 6.Compare data and make decision
	if slices.Compare(msgIDs1, msgIDs2) != 0 {
		err = fmt.Errorf("message IDs are not the same: from insertion %v, from getting %v", msgIDs1, msgIDs2)
		return
	}

	return
}
