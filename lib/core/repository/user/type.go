package userrepo

import (
	"context"

	"github.com/cangkir13/tes-go-user/lib/core/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

var _ UserRepositoryInterface = &UserRepository{}

type UserRepositoryInterface interface {
	GetList(ctx context.Context) (output []*model.User, err error)
	GetOneByID(ctx context.Context, userID int) (output *model.User, err error)
	CreateOne(ctx context.Context, props CreateOneProps) (newId int, err error)
	// UpdateOne(ctx context.Context, userID int, props UpdateOneProps) error
	// DeleteOne(ctx context.Context, userID int) error
}
