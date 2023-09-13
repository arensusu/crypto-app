package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"

	"funding-rate/coinglass"
	"funding-rate/database"
	"funding-rate/pair"
	"funding-rate/telegram"
	"funding-rate/user"
	"funding-rate/watchlist"
)

func main() {
	db := database.LoadDatabase()
	api := coinglass.NewCoinglassApi(coinglass.ApiEndpoint, os.Getenv("COINGLASS_APIKEY"))

	userRepo := user.NewUserPostgresRepository(db)
	watchlistRepo := watchlist.NewWatchlistPostgresRepository(db)
	fundingRepo := pair.NewPairPostgresRepository(db, &api)

	userUsecase := user.NewUserUsecase(userRepo)
	watchlistUsecase := watchlist.NewWatchlistUsecase(watchlistRepo)
	fundingUsecase := pair.NewPairUsecase(watchlistRepo, fundingRepo)

	// tgbot
	tgbot := telegram.NewTelegramBot()
	telegramHandler := telegram.NewTelegramHandler(tgbot, userUsecase, watchlistUsecase, fundingUsecase)

	go telegramHandler.Run()
}
