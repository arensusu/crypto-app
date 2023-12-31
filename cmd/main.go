package main

import (
	"log"
	"net/http"

	"github.com/arensusu/crypto-app/api"
	"github.com/arensusu/crypto-app/exchange/binance"
	"github.com/arensusu/crypto-app/exchange/binance_future"
	"github.com/arensusu/crypto-app/exchange/bitget"
	"github.com/arensusu/crypto-app/exchange/bybit"
	"github.com/arensusu/crypto-app/pkg/assets"
	"github.com/arensusu/crypto-app/pkg/crossexchange"

	"github.com/gin-gonic/gin"
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
	binance_future := binance_future.New()
	bitget := bitget.New()
	ss := crossexchange.NewCrossExchangeSingleSymbol(bybit, binance_future, bitget)
	asset := assets.NewAssetsFinder(bybit, binance, binance_future, bitget)

	router := gin.Default()
	api.NewCrossExchangeServer(router, ss)
	api.NewAssetsServer(router, asset)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
	// assetsUsecase := assets.NewAssetsUsecase([]any{bybit, binance, binance_future, bitget})
	// assetsUsecase.GetAssets()
}
