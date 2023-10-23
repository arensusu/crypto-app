package user

import (
	"crypto-exchange/domain"
	"fmt"

	"gorm.io/gorm"
)

type UserPostgresRepository struct {
	db *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) domain.UserRepository {
	return &UserPostgresRepository{db}
}

func (repo *UserPostgresRepository) CreateUser(chatID int64) error {
	user := domain.User{ChatID: chatID}
	err := repo.db.Create(&user).Error
	if err != nil {
		return fmt.Errorf("create user failed: %w", err)
	}
	return nil
}

func (repo *UserPostgresRepository) RetrieveUsers() ([]domain.User, error) {
	var users []domain.User
	err := repo.db.Find(&users).Error
	if err != nil {
		return []domain.User{}, fmt.Errorf("get user failed: %w", err)
	}
	return users, nil
}
