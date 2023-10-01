package main

import (
	"errors"
	"fmt"
	"funding-rate/exchange"
	binanceEx "funding-rate/exchange/binance"
	bitgetEx "funding-rate/exchange/bitget"
	bybitEx "funding-rate/exchange/bybit"
	"funding-rate/exchange/strategy"
	"log"
	"os"
	"sync"
	"time"

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
	binance := binanceEx.New(os.Getenv("BINANCE_API_KEY"), os.Getenv("BINANCE_API_SECRET"))
	bitget := bitgetEx.New()

	exchange := exchange.New(binance, bybit, bitget)
	wg := new(sync.WaitGroup)
	wg.Add(len(exchange.List))
	start := time.Now().UnixMilli()
	for _, ex := range exchange.List {
		go func(ex any) {
			defer wg.Done()
			strat, ok := ex.(strategy.CrossExArbitrage)
			if !ok {
				log.Fatal(errors.New("type error"))
			}
			res, err := strat.GetCrossExArbitrageResponse("OGN")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%+v\n", res)
		}(ex)
	}
	wg.Wait()
	fmt.Println(time.Now().UnixMilli() - start)
}
