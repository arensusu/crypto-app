package main

import (
	binance "funding-rate/exchange/binance"
	bitget "funding-rate/exchange/bitget"
	bybit "funding-rate/exchange/bybit"
	"funding-rate/exchange/strategy"

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
	bybit := bybit.New()
	binance := binance.New()
	bitget := bitget.New()

	exchange := strategy.New(binance, bybit, bitget)
	exchange.GetSingleCrossExchangeArbitrage("OGN")
}
