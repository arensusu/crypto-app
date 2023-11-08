package main

import (
	"crypto-exchange/api"
	"crypto-exchange/exchange/binance"
	"crypto-exchange/exchange/binance_future"
	"crypto-exchange/exchange/bitget"
	"crypto-exchange/exchange/bybit"
	"crypto-exchange/pkg/crossexchange"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	_ = binance.New()
	binance_future := binance_future.New()
	bitget := bitget.New()
	ss := crossexchange.NewCrossExchangeSingleSymbol(bybit, binance_future, bitget)

	router := mux.NewRouter()
	api.NewCrossExchangeServer(router, ss)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
	// assetsUsecase := assets.NewAssetsUsecase([]any{bybit, binance, binance_future, bitget})
	// assetsUsecase.GetAssets()
}
