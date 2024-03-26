package usecase

import (
	"context"

	"github.com/mahdifr17/sirka-test/entity"
)

type UsecaseInterface interface {
	DisplayUser(ctx context.Context, userID string) (entity.User, error)
	DisplayAllUsers(ctx context.Context) ([]entity.User, error)
}
