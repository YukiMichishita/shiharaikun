package repository

import (
	"context"
	"shiharaikun/internal/domain/model"
)

type UserRepository interface {
	GetBySessionID(ctx context.Context, sessionID string) (*model.User, error)
}
