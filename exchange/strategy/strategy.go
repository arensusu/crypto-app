package strategy

import (
	"fmt"
	"log"
	"strconv"
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
	fmt.Println(result["STRAXUSDT"])

	results := SymbolFundingPriceDiffs{}
	for symbol, fundingPrices := range result {
		diffs := []FundingPriceDiff{}
		for i := 0; i < len(fundingPrices); i += 1 {
			for j := i + 1; j < len(fundingPrices); j += 1 {
				diffs = append(diffs, FundingPriceDiff{
					ExchangeBuy:     fundingPrices[j].Exchange,
					ExchangeSell:    fundingPrices[i].Exchange,
					PriceDiff:       rounding(fundingPrices[j].Price/fundingPrices[i].Price - 1.0),
					FundingRateDiff: rounding(fundingPrices[j].FundingRate - fundingPrices[i].FundingRate),
					FundingTime:     time.UnixMilli(fundingPrices[j].FundingTime).String() + "/" + time.UnixMilli(fundingPrices[i].FundingTime).String(),
				})
			}
		}
		results[symbol] = diffs
	}
	fmt.Println(results["STRAXUSDT"])
}

type FundingPrice struct {
	Exchange    string
	Price       float64
	FundingRate float64
	FundingTime int64
}

type SymbolExchangeFundingPrice map[string][]FundingPrice

func (m SymbolExchangeFundingPrice) Set(exchange, symbol, price, fundingRate, fundingTime string) {
	if _, exist := m[symbol]; !exist {
		m[symbol] = []FundingPrice{}
	}

	priceFloat64, _ := strconv.ParseFloat(price, 64)
	fundingRateFloat64, _ := strconv.ParseFloat(fundingRate, 64)
	fundingTimeInt64, _ := strconv.ParseInt(fundingTime, 10, 64)

	m[symbol] = append(m[symbol], FundingPrice{
		Exchange:    exchange,
		Price:       priceFloat64,
		FundingRate: fundingRateFloat64,
		FundingTime: fundingTimeInt64,
	})
}

func (m SymbolExchangeFundingPrice) SetSpecial(exchange, symbol, fundingRate, fundingTime string) {
	fundingRateFloat64, _ := strconv.ParseFloat(fundingRate, 64)
	fundingTimeInt64, _ := strconv.ParseInt(fundingTime, 10, 64)
	if elem, ok := m[symbol]; ok {
		for i, fundingPrice := range elem {
			if fundingPrice.Exchange == exchange {
				m[symbol][i].FundingRate = fundingRateFloat64
				m[symbol][i].FundingTime = fundingTimeInt64
				break
			}
		}
	}
}

func rounding(x float64) float64 {
	if x < 0.0001 && x > -0.0001 {
		return 0.0
	}
	return x
}
