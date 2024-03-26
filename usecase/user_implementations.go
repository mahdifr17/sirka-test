package usecase

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/mahdifr17/sirka-test/entity"
	"github.com/mahdifr17/sirka-test/repository"
)

type UseCase struct {
	Repository repository.RepositoryInterface
}

func NewUseCase(repository repository.RepositoryInterface) *UseCase {
	return &UseCase{Repository: repository}
}

func (u *UseCase) DisplayUser(ctx context.Context, userID string) (entity.User, error) {
	usr, err := u.Repository.GetUser(ctx, userID)
	if err != nil {
		return entity.User{}, err
	}
	return usr, nil
}

func (u *UseCase) DisplayAllUsers(ctx context.Context) ([]entity.User, error) {
	listUsr, err := u.Repository.GetAllUsers(ctx)
	if err != nil {
		log.Error("UseCase.DisplayAllUsers()", err)
	}
	return listUsr, err
}
