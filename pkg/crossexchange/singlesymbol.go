package crossexchange

import (
	"crypto-exchange/exchange/domain"
	"fmt"
	"math"
	"sync"
	"time"
)

type CrossExchangeSingleSymbol struct {
	Exchanges []any
}

func NewCrossExchangeSingleSymbol(exchanges []any) *CrossExchangeSingleSymbol {
	return &CrossExchangeSingleSymbol{Exchanges: exchanges}
}

type SingleSymbolResult struct {
	Data   []domain.FundingPrice
	Diff   []FundingPriceDiff
	Errors []error
}

type FundingPriceDiff struct {
	ExchangeBuy     string
	ExchangeSell    string
	PriceDiff       float64
	FundingRateDiff float64
	FundingTime     string
}

func (c *CrossExchangeSingleSymbol) Do(symbol string) SingleSymbolResult {
	fmt.Println(symbol)

	raw, errs := c.getExchangeData(symbol)

	processed := calculateFundingPrices(raw)

	return SingleSymbolResult{Data: raw, Diff: processed, Errors: errs}
}

func (c *CrossExchangeSingleSymbol) getExchangeData(symbol string) ([]domain.FundingPrice, []error) {
	wg := new(sync.WaitGroup)
	wg.Add(len(c.Exchanges))

	dataLock := new(sync.Mutex)
	errsLock := new(sync.Mutex)
	data := []domain.FundingPrice{}
	errs := []error{}
	for _, ex := range c.Exchanges {
		go func(ex any) {
			defer wg.Done()
			strat, ok := ex.(domain.GetFundingAndPricer)
			if !ok {
				errsLock.Lock()
				errs = append(errs, fmt.Errorf("does not have this function"))
				return
			}

			info, err := (strat).GetFundingAndPrice(symbol)
			if err != nil {
				errsLock.Lock()
				errs = append(errs, err)
				errsLock.Unlock()
				return
			}
			dataLock.Lock()
			data = append(data, *info)
			dataLock.Unlock()
		}(ex)
	}
	wg.Wait()

	return data, errs
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
