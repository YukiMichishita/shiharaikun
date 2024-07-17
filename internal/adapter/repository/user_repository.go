package repository

import (
	"context"
	"fmt"
	"shiharaikun/internal/adapter/db/query"
	"shiharaikun/internal/domain/model"
	"shiharaikun/internal/domain/repository"
)

type userRepository struct{}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (u *userRepository) GetBySessionID(ctx context.Context, sessionID string) (*model.User, error) {
	uq := query.Q.User
	user, err := uq.WithContext(ctx).Preload(uq.Company).Where(uq.SessionID.Eq(sessionID)).First()
	if err != nil {
		return nil, fmt.Errorf("failed to get user by session id: %w", err)
	}
	return user, nil
}
