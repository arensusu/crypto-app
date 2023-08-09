package domain

type User struct {
	ChatID int64 `gorm:"primarykey"`
}

type Notification struct {
	ChatID  int64
	Message string
}

type IUserRepository interface {
	AddUser(chatID int64) error
	RetrieveUsers() ([]User, error)
}

type IUserUseCase interface {
	NewUser(chatID int64) string
	GetUsersNotification() []Notification
}
