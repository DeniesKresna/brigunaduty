package usecase

import "github.com/DeniesKresna/brigunaduty/service/modules/user/repository"

type UserUsecase struct {
	userRepo repository.UserRepository
}

func UserCreateUsecase(userRepo repository.UserRepository) UserUsecase {
	userUsecase := UserUsecase{
		userRepo: userRepo,
	}
	return userUsecase
}
