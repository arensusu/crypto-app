package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"

	"funding-rate/coinglass"
	"funding-rate/database"
	"funding-rate/funding"
	httpapi "funding-rate/http"
	"funding-rate/telegram"
	"funding-rate/user"
	"funding-rate/watchlist"
)

func main() {
	db := database.LoadDatabase()
	api := coinglass.NewCoinglassApi(coinglass.ApiEndpoint, os.Getenv("COINGLASS_APIKEY"))

	userRepo := user.NewUserPostgresRepository(db)
	watchlistRepo := watchlist.NewWatchlistPostgresRepository(db)
	fundingRepo := funding.NewFundingPostgresRepository(db, &api)

	userUsecase := user.NewUserUsecase(userRepo)
	watchlistUsecase := watchlist.NewWatchlistUsecase(watchlistRepo)
	fundingUsecase := funding.NewFundingUsecase(watchlistRepo, fundingRepo)

	// tgbot
	tgbot := telegram.NewTelegramBot()
	telegramHandler := telegram.NewTelegramHandler(tgbot, userUsecase, watchlistUsecase, fundingUsecase)

	go telegramHandler.Run()

	// restful api
	router := mux.NewRouter()
	httpapi.NewFundingHandler(router, fundingUsecase)
	http.ListenAndServe(":8000", router)
}
