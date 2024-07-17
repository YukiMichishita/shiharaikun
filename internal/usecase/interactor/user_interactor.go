package interactor

import (
	"context"
	"fmt"
	"shiharaikun/internal/domain/model"
	"shiharaikun/internal/domain/repository"
)

type UserInterActor struct {
	repo repository.UserRepository
}

func NewUserInterActor(repo repository.UserRepository) *UserInterActor {
	return &UserInterActor{repo: repo}
}

func (u *UserInterActor) GetUserBySessionID(ctx context.Context, sessionID string) (*model.User, error) {
	user, err := u.repo.GetBySessionID(ctx, sessionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by session id: %w", err)
	}
	return user, nil
}
