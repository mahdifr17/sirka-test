package repository

import (
	"context"

	"github.com/mahdifr17/sirka-test/entity"
)

type RepositoryInterface interface {
	GetUser(ctx context.Context, userID string) (usr entity.User, err error)
	GetAllUsers(ctx context.Context) (listUsr []entity.User, err error)
}
