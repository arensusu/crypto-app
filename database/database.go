package database

import (
	"crypto-exchange/domain"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	hostname = os.Getenv("POSTGRES_HOSTNAME")
	user     = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname   = os.Getenv("POSTGRES_DBNAME")
	port     = os.Getenv("POSTGRES_PORT")
	sslmode  = "disable"
)

func LoadDatabase() *gorm.DB {
	db := Connect()
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Watchlist{})
	db.AutoMigrate(&domain.PerpetualWatchlist{})
	db.AutoMigrate(&domain.FundingSearched{})
	db.AutoMigrate(&domain.PrevPrice{})
	return db
}

func Connect() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Taipei",
		hostname, port, user, password, dbname, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
