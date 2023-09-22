package userrepo

import (
	"context"
	"errors"
	"time"

	"github.com/cangkir13/tes-go-user/lib/core/model"
	"gorm.io/gorm"
)

func (repo *UserRepository) GetList(ctx context.Context) (output []*model.User, err error) {
	query := repo.DB.WithContext(ctx)
	if err = query.Find(&output).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) GetOneByID(ctx context.Context, userID int) (output *model.User, err error) {
	query := repo.DB.WithContext(ctx)

	if err = query.Where("id = ?", userID).First(&output).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.New("data not found")
			return
		}
		return
	}
	return
}

type UserFieldsData struct {
	Name      string
	Email     string
	Password  *string
	RoleID    int
	Status    *int
	UpdatedAt *time.Time
}

func (fields *UserFieldsData) toModel() model.User {
	return model.User{
		Name:      fields.Name,
		Email:     fields.Email,
		Password:  *fields.Password,
		RoleID:    fields.RoleID,
		UpdatedAt: *fields.UpdatedAt,
	}
}

func (fields *UserFieldsData) collumnsUpdated() []string {
	return []string{
		"name",
		"email",
		"password",
		"role_id",
		"status",
		"updated_at",
	}
}

type CreateOneProps struct {
	Data UserFieldsData
}

func (repo *UserRepository) CreateOne(ctx context.Context, props CreateOneProps) (newId int, err error) {
	query := repo.DB

	var newData = props.Data.toModel()
	if err = query.WithContext(ctx).Create(&newData).Error; err != nil {
		newId = 0
		return
	}

	newId = newData.ID
	return
}

type UpdateOneProps struct {
	Data UserFieldsData
}

func (repo *UserRepository) UpdateOne(ctx context.Context, userID int, props UpdateOneProps) error {

	var updateuser = props.Data.toModel()

	query := repo.DB.WithContext(ctx).
		Where("id = ?", userID).
		Select(props.Data.collumnsUpdated())
	if err := query.Updates(&updateuser).Error; err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) DeleteOne(ctx context.Context, userID int) error {
	query := repo.DB.WithContext(ctx)
	if err := query.Where("id = ?", userID).Delete(&model.User{}).Error; err != nil {
		return err
	}

	return nil
}
