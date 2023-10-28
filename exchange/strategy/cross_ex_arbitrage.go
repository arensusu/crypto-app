package strategy

import (
	"fmt"
	"log"
	"math"
	"sync"
	"time"
)

type CrossExArbitrageInformation struct {
	ExchangeName    string
	Symbol          string
	LastPrice       string
	FundingRate     string
	NextFundingTime string
}

type CrossExArbitrager interface {
	GetCrossExArbitrageInformation() (SymbolExchangeFundingPrice, error)
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
	result := SymbolExchangeFundingPrice{}
	_ = result
	for _, ex := range exs.Exchanges {
		go func(ex any) {
			defer wg.Done()
			strat, ok := ex.(CrossExArbitrager)
			if !ok {
				log.Fatal(fmt.Errorf("type error: %v", ex))
			}

			info, err := (strat).GetCrossExArbitrageInformation()
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
	fmt.Println(result["TRBUSDT"])

	results := SymbolFundingPriceDiffs{}
	for symbol, fundingPrices := range result {

		results[symbol] = calculateFundingPrices(fundingPrices)
	}
	fmt.Println(results["TRBUSDT"])
}

func calculateFundingPrices(fps []FundingPrice) []FundingPriceDiff {
	diffs := []FundingPriceDiff{}
	for i := 0; i < len(fps); i += 1 {
		for j := i + 1; j < len(fps); j += 1 {
			diffs = append(diffs, getBestCrossExchangeBuySell(&fps[i], &fps[j]))
		}
	}
	return diffs
}

func getBestCrossExchangeBuySell(fpA, fpB *FundingPrice) FundingPriceDiff {
	var buy, sell *FundingPrice

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

func isFundingTimeNotTheSame(fpA, fpB *FundingPrice) bool {
	return fpA.FundingTime != fpB.FundingTime && fpA.FundingTime != 0 && fpB.FundingTime != 0
}

func getBuySellByFundingTime(fpA, fpB *FundingPrice) (*FundingPrice, *FundingPrice) {
	if fpA.FundingTime <= fpB.FundingTime {
		return fpA, fpB
	}
	return fpB, fpA
}

func isPriceDiffLarger(fpA, fpB *FundingPrice) bool {
	priceDiff := fpA.Price/fpB.Price - 1
	fundingRateDiff := fpA.FundingRate - fpB.FundingRate
	return math.Abs(priceDiff) > math.Abs(fundingRateDiff)
}

func getBuySellByPrice(fpA, fpB *FundingPrice) (*FundingPrice, *FundingPrice) {
	if fpA.Price <= fpB.Price {
		return fpA, fpB
	}
	return fpB, fpA
}

func getBuySellByFundingRate(fpA, fpB *FundingPrice) (*FundingPrice, *FundingPrice) {
	if fpA.FundingRate <= fpB.FundingRate {
		return fpA, fpB
	}
	return fpB, fpA
}
