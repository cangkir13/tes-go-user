package userusecase

import (
	"context"

	"github.com/cangkir13/tes-go-user/lib/core/model"
	userrepo "github.com/cangkir13/tes-go-user/lib/core/repository/user"
)

type UserUsecase struct {
	userrepo userrepo.UserRepositoryInterface
}

var _ UserUsecaseInterface = &UserUsecase{}

type UserUsecaseInterface interface {
	GetListUser(ctx context.Context) ([]*model.User, error)
}
