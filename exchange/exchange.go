package exchange

import (
	"errors"
	"fmt"
	"funding-rate/exchange/strategy"
	"log"
	"sort"
	"sync"
	"time"
)

type Exchanges []interface{}

func New(exchanges ...interface{}) *Exchanges {
	return &Exchanges{exchanges}
}

func (exs *Exchanges) GetSingleCrossExchangeArbitrage(coin string) {
	start := time.Now().UnixMilli()

	wg := new(sync.WaitGroup)
	wg.Add(len(*exs))

	lock := new(sync.Mutex)
	infos := []*strategy.CrossExArbitrageInformation{}

	for _, ex := range *exs {
		go func(ex any) {
			defer wg.Done()
			strat, ok := ex.(strategy.CrossExArbitrager)
			if !ok {
				log.Fatal(errors.New("type error"))
			}

			info, err := strat.GetCrossExArbitrageInformation("OGN")
			if err != nil {
				log.Fatal(err)
			}

			lock.Lock()
			infos = append(infos, info)
			lock.Unlock()
		}(ex)
	}
	wg.Wait()

	sort.Slice(infos, func(i, j int) bool {
		return infos[i].LastPrice < infos[j].LastPrice
	})

	results := []*strategy.CrossExArbitrageResult{}
	for i := 0; i < len(infos); i += 1 {
		for j := i + 1; j < len(infos); j += 1 {
			results = append(results, &strategy.CrossExArbitrageResult{
				ExchangePair:     infos[j].ExchangeName + "/" + infos[i].ExchangeName,
				PriceDiffPercent: infos[j].LastPrice/infos[i].LastPrice - 1.0,
				FundingRateDiff:  infos[j].FundingRate - infos[i].FundingRate,
				NextFundingTime:  time.UnixMilli(infos[j].NextFundingTime).String() + "/" + time.UnixMilli(infos[i].NextFundingTime).String(),
			})
		}
	}
	fmt.Println(time.Now().UnixMilli() - start)
	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}
}
