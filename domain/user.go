package domain

type User struct {
	ChatID int64 `gorm:"primarykey"`
}

type UserRepository interface {
	CreateUser(chatID int64) error
	RetrieveUsers() ([]User, error)
}

type UserUsecase interface {
	AddUser(chatID int64) error
	GetUsers() ([]User, error)
}
