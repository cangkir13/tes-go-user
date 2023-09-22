package userusecase

import (
	"context"
	"time"

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
			Name:      props.Name,
			Email:     props.Email,
			Password:  &props.Password,
			RoleID:    props.RoleID,
			Status:    &props.Status,
			UpdatedAt: &props.CreatedAt,
		},
	})

	if err != nil {
		return 0, err
	}

	return data, err
}

func (u *UserUsecase) UpdateOne(ctx context.Context, userID int, props model.User) error {
	// find id user
	_, err := u.userrepo.GetOneByID(ctx, userID)

	if err != nil {
		return err
	}

	dateTime := time.Now()

	err = u.userrepo.UpdateOne(ctx, userID, userrepo.UpdateOneProps{
		Data: userrepo.UserFieldsData{
			Name:      props.Name,
			Password:  &props.Password,
			RoleID:    2,
			Status:    &props.Status,
			UpdatedAt: &dateTime,
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) DeleteOne(ctx context.Context, userID int) error {
	// find id user
	_, err := u.userrepo.GetOneByID(ctx, userID)

	if err != nil {
		return err
	}

	// delete process
	err = u.userrepo.DeleteOne(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
