package repository

import (
	"context"
	"errors"

	"github.com/mahdifr17/sirka-test/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

type NewRepositoryOptions struct {
	Dsn string
}

func NewRepository(opts NewRepositoryOptions) *Repository {
	db, err := gorm.Open(
		postgres.Open(opts.Dsn),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}
	return &Repository{db}
}

var (
	errPhoneNumberUnique error = errors.New("error phone number already exist")
)

func (r *Repository) GetUser(ctx context.Context, userID string) (usr entity.User, err error) {
	if err := r.Db.Find(&usr, "user_id = ?", userID).Error; err != nil {
		return entity.User{}, err
	}

	return usr, nil
}

func (r *Repository) GetAllUsers(ctx context.Context) (listUsr []entity.User, err error) {
	if err := r.Db.Find(&listUsr).Error; err != nil {
		return nil, err
	}
	return listUsr, nil
}
