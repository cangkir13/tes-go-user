package userusecase

import (
	"context"

	"github.com/cangkir13/tes-go-user/lib/core/model"
)

func (u *UserUsecase) GetListUser(ctx context.Context) ([]*model.User, error) {
	data, err := u.userrepo.GetList(ctx)

	if err != nil {
		return nil, err
	}

	return data, err
}
