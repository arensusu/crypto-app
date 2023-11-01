package cross

import (
	"crypto-exchange/domain"
	"fmt"
	"log"
	"math"
	"sync"
	"time"
)

type StrategyExecuter struct {
	Exchanges []interface{}
}

func New(exchanges ...interface{}) *StrategyExecuter {
	ex := &StrategyExecuter{Exchanges: exchanges}
	return ex
}

type CrossExArbitrageInformation struct {
	ExchangeName    string
	Symbol          string
	LastPrice       string
	FundingRate     string
	NextFundingTime string
}

type FundingPriceDiff struct {
	ExchangeBuy     string
	ExchangeSell    string
	PriceDiff       float64
	FundingRateDiff float64
	FundingTime     string
}

type SymbolFundingPriceDiffs map[string][]FundingPriceDiff

func (exs *StrategyExecuter) GetCrossExchangeArbitrage() {
	//start := time.Now().UnixMilli()

	wg := new(sync.WaitGroup)
	wg.Add(len(exs.Exchanges))

	lock := new(sync.Mutex)
	result := domain.FundingPricesOfSymbol{}
	_ = result
	for _, ex := range exs.Exchanges {
		go func(ex any) {
			defer wg.Done()
			strat, ok := ex.(domain.GetFundingAndPricer)
			if !ok {
				log.Fatal(fmt.Errorf("type error: %v", ex))
			}

			info, err := (strat).GetFundingAndPrices()
			if err != nil {
				log.Fatal(err)
			}
			lock.Lock()
			for k, v := range info {
				result[k] = append(result[k], v...)
			}
			lock.Unlock()
		}(ex)
	}
	wg.Wait()

	results := SymbolFundingPriceDiffs{}
	for symbol, fundingPrices := range result {

		results[symbol] = calculateFundingPrices(fundingPrices)
	}
	fmt.Println(results["POLYXUSDT"])
}

func calculateFundingPrices(fps []domain.FundingPrice) []FundingPriceDiff {
	diffs := []FundingPriceDiff{}
	for i := 0; i < len(fps); i += 1 {
		for j := i + 1; j < len(fps); j += 1 {
			diffs = append(diffs, getBestCrossExchangeBuySell(&fps[i], &fps[j]))
		}
	}
	return diffs
}

func getBestCrossExchangeBuySell(fpA, fpB *domain.FundingPrice) FundingPriceDiff {
	var buy, sell *domain.FundingPrice

	if isFundingTimeNotTheSame(fpA, fpB) {
		buy, sell = getBuySellByFundingTime(fpA, fpB)
	} else if isPriceDiffLarger(fpA, fpB) {
		buy, sell = getBuySellByPrice(fpA, fpB)
	} else {
		buy, sell = getBuySellByFundingRate(fpA, fpB)
	}

	diff := FundingPriceDiff{
		ExchangeBuy:     buy.Exchange,
		ExchangeSell:    sell.Exchange,
		PriceDiff:       buy.Price / sell.Price,
		FundingRateDiff: buy.FundingRate - sell.FundingRate,
		FundingTime:     time.UnixMilli(buy.FundingTime).In(time.FixedZone("UTC+8", 8*60*60)).Format("2006/01/02 15:04 MST"),
	}

	return diff
}

func isFundingTimeNotTheSame(fpA, fpB *domain.FundingPrice) bool {
	return fpA.FundingTime != fpB.FundingTime && fpA.FundingTime != 0 && fpB.FundingTime != 0
}

func getBuySellByFundingTime(fpA, fpB *domain.FundingPrice) (*domain.FundingPrice, *domain.FundingPrice) {
	if fpA.FundingTime <= fpB.FundingTime {
		return fpA, fpB
	}
	return fpB, fpA
}

func isPriceDiffLarger(fpA, fpB *domain.FundingPrice) bool {
	priceDiff := fpA.Price/fpB.Price - 1
	fundingRateDiff := fpA.FundingRate - fpB.FundingRate
	return math.Abs(priceDiff) > math.Abs(fundingRateDiff)
}

func getBuySellByPrice(fpA, fpB *domain.FundingPrice) (*domain.FundingPrice, *domain.FundingPrice) {
	if fpA.Price <= fpB.Price {
		return fpA, fpB
	}
	return fpB, fpA
}

func getBuySellByFundingRate(fpA, fpB *domain.FundingPrice) (*domain.FundingPrice, *domain.FundingPrice) {
	if fpA.FundingRate <= fpB.FundingRate {
		return fpA, fpB
	}
	return fpB, fpA
}
