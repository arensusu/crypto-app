package main

import (
	"funding-rate/exchange"
	binanceEx "funding-rate/exchange/binance"
	bitgetEx "funding-rate/exchange/bitget"
	bybitEx "funding-rate/exchange/bybit"

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
	bybit := bybitEx.New()
	binance := binanceEx.New()
	bitget := bitgetEx.New()

	exchange := exchange.New(binance, bybit, bitget)
	exchange.GetSingleCrossExchangeArbitrage("ETH")
}
