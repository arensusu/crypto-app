package main

import (
	"errors"
	"fmt"
	"funding-rate/exchange"
	bybitEx "funding-rate/exchange/bybit"
	"funding-rate/exchange/strategy"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// db := database.LoadDatabase()
	// api := coinglass.NewCoinglassApi(coinglass.ApiEndpoint, os.Getenv("COINGLASS_APIKEY"))

	// userRepo := user.NewUserPostgresRepository(db)
	// watchlistRepo := watchlist.NewWatchlistPostgresRepository(db)
	// fundingRepo := pair.NewPairPostgresRepository(db, &api)

	// userUsecase := user.NewUserUsecase(userRepo)
	// watchlistUsecase := watchlist.NewWatchlistUsecase(watchlistRepo)
	// fundingUsecase := pair.NewPairUsecase(watchlistRepo, fundingRepo)

	// // tgbot
	// tgbot := telegram.NewTelegramBot()
	// telegramHandler := telegram.NewTelegramHandler(tgbot, userUsecase, watchlistUsecase, fundingUsecase)

	// go telegramHandler.Run()
	bybit := bybitEx.New(os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))
	exchange := exchange.New(bybit)
	for _, ex := range exchange.List {
		strat, ok := ex.(strategy.CrossExArbitrage)
		if !ok {
			log.Fatal(errors.New("type error"))
		}
		res, err := strat.GetCrossExArbitrageResponse("ETH")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", res)
	}

}
