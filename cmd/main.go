package main

import (
	"crypto-exchange/exchange/binance"
	"crypto-exchange/exchange/binance_future"
	"crypto-exchange/exchange/bitget"
	"crypto-exchange/exchange/bybit"
	"flag"
	"fmt"

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
	doCross := flag.String("cross", "", "")
	flag.Parse()

	bybit := bybit.New()
	binance := binance.New()
	binance_future := binance_future.New()
	bitget := bitget.New()
	_, _, _ = binance, bybit, binance_future
	exchange := NewExchanges(binance_future, bybit, bitget)

	if *doCross != "" {
		fmt.Println(exchange.GetCrossExchangeFundingPrice(*doCross))
	}

	// assetsUsecase := assets.NewAssetsUsecase([]any{bybit, binance, binance_future, bitget})
	// assetsUsecase.GetAssets()
}
