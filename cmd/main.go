package main

import (
	_ "github.com/joho/godotenv/autoload"

	"funding-rate/database"
	"funding-rate/funding"
	"funding-rate/telegram"
	"funding-rate/user"
)

func main() {
	db := database.LoadDatabase()
	fundingRepo := funding.NewPostgresFundingRepository(db)
	userRepo := user.NewPostgresUserRepository(db)

	fundingUseCase := funding.NewFundingUseCase(fundingRepo)
	userUseCase := user.NewUserUseCase(userRepo, fundingRepo)

	telegramHandler := telegram.NewTelegramHandler(fundingUseCase, userUseCase)

	telegramHandler.Start()

}
