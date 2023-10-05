package strategy

import (
	"errors"
	"fmt"
	"log"
	"sort"
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

func (exs *StrategyExecuter) GetSingleCrossExchangeArbitrage(coin string) {
	start := time.Now().UnixMilli()

	wg := new(sync.WaitGroup)
	wg.Add(len(exs.Exchanges))

	lock := new(sync.Mutex)
	infos := []*CrossExArbitrageInformation{}

	for _, ex := range exs.Exchanges {
		go func(ex any) {
			defer wg.Done()
			strat, ok := ex.(CrossExArbitrager)
			if !ok {
				log.Fatal(errors.New("type error"))
			}

			info, err := (strat).GetCrossExArbitrageInformation(coin)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(info)
			lock.Lock()
			infos = append(infos, info)
			lock.Unlock()
		}(ex)
	}
	wg.Wait()

	sort.Slice(infos, func(i, j int) bool {
		return infos[i].LastPrice < infos[j].LastPrice
	})

	results := []*CrossExArbitrageResult{}
	for i := 0; i < len(infos); i += 1 {
		for j := i + 1; j < len(infos); j += 1 {
			results = append(results, &CrossExArbitrageResult{
				ExchangePair:     infos[j].ExchangeName + "/" + infos[i].ExchangeName,
				PriceDiffPercent: rounding(infos[j].LastPrice/infos[i].LastPrice - 1.0),
				FundingRateDiff:  rounding(infos[j].FundingRate - infos[i].FundingRate),
				NextFundingTime:  time.UnixMilli(infos[j].NextFundingTime).String() + "/" + time.UnixMilli(infos[i].NextFundingTime).String(),
			})
		}
	}
	fmt.Println(time.Now().UnixMilli() - start)
	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}
}

func rounding(x float64) float64 {
	if x < 0.001 {
		return 0.0
	}
	return x
}
