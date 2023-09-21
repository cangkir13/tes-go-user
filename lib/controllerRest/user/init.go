package usercontroller

import (
	userusecase "github.com/cangkir13/tes-go-user/lib/core/usecase/user"
)

type UserController struct {
	userusecase userusecase.UserUsecaseInterface
}

func NewUserController(
	userusecase userusecase.UserUsecaseInterface,
) *UserController {
	return &UserController{userusecase: userusecase}
}
