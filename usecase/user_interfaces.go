package usecase

import (
	"context"

	"github.com/sample-tech-test/entity"
)

type UsecaseInterface interface {
	DisplayUser(ctx context.Context, userID string) (entity.User, error)
	DisplayAllUsers(ctx context.Context) ([]entity.User, error)
}
