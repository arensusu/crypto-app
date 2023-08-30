package user

import (
	"funding-rate/domain"
)

type UserUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{userRepo}
}

func (usecase *UserUsecase) AddUser(chatID int64) error {
	err := usecase.userRepo.CreateUser(chatID)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *UserUsecase) GetUsers() ([]domain.User, error) {
	users, err := usecase.userRepo.RetrieveUsers()
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}
