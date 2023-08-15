package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"

	"funding-rate/coinglass"
	"funding-rate/database"
	"funding-rate/funding"
	"funding-rate/telegram"
	"funding-rate/user"
)

func main() {
	db := database.LoadDatabase()
	api := coinglass.NewCoinglassApi(coinglass.ApiEndpoint, os.Getenv("COINGLASS_APIKEY"))

	fundingRepo := funding.NewPostgresFundingRepository(db, &api)
	userRepo := user.NewPostgresUserRepository(db)

	fundingUseCase := funding.NewFundingUseCase(fundingRepo)
	userUseCase := user.NewUserUseCase(userRepo, fundingRepo)

	telegramHandler := telegram.NewTelegramHandler(fundingUseCase, userUseCase)

	telegramHandler.Start()

}
