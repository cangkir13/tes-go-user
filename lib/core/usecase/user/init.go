package userusecase

import userrepo "github.com/cangkir13/tes-go-user/lib/core/repository/user"

func NewUserUsecase(
	userrepo userrepo.UserRepositoryInterface,
) *UserUsecase {
	return &UserUsecase{userrepo: userrepo}
}
