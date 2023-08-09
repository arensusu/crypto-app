package user

import (
	"funding-rate/domain"

	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) domain.IUserRepository {
	return &PostgresUserRepository{db}
}

func (repo *PostgresUserRepository) AddUser(chatID int64) error {
	user := domain.User{ChatID: chatID}
	err := repo.db.Create(&user).Error
	return err
}

func (repo *PostgresUserRepository) RetrieveUsers() ([]domain.User, error) {
	var users []domain.User
	err := repo.db.Find(&users).Error
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}
