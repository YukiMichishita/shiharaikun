package usecase

import (
	"context"
	"shiharaikun/internal/domain/model"
)

type UserUseCase interface {
	GetUserBySessionID(ctx context.Context, sessionID string) (*model.User, error)
}
