package userusecase

import (
	"context"

	"github.com/cangkir13/tes-go-user/lib/core/model"
	userrepo "github.com/cangkir13/tes-go-user/lib/core/repository/user"
)

func (u *UserUsecase) GetListUser(ctx context.Context) ([]*model.User, error) {
	data, err := u.userrepo.GetList(ctx)

	if err != nil {
		return nil, err
	}

	return data, err
}

func (u *UserUsecase) GetOneByID(ctx context.Context, userID int) (*model.User, error) {
	data, err := u.userrepo.GetOneByID(ctx, userID)

	if err != nil {
		return nil, err
	}

	return data, err
}

func (u *UserUsecase) CreateOne(ctx context.Context, props model.User) (int, error) {
	data, err := u.userrepo.CreateOne(ctx, userrepo.CreateOneProps{
		Data: userrepo.UserFieldsData{
			Name:     &props.Name,
			Email:    &props.Email,
			Password: &props.Password,
			RoleID:   &props.RoleID,
			Status:   &props.Status,
		},
	})

	if err != nil {
		return 0, err
	}

	return data, err
}
