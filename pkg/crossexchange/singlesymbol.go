package crossexchange

import (
	"sync"

	"github.com/arensusu/crypto-app/exchange"
	"github.com/arensusu/crypto-app/exchange/types"
)

type CrossExchangeSingleSymbol struct {
	Exchanges []exchange.Exchange
}

func NewCrossExchangeSingleSymbol(exchanges ...exchange.Exchange) *CrossExchangeSingleSymbol {
	return &CrossExchangeSingleSymbol{Exchanges: exchanges}
}

func (c *CrossExchangeSingleSymbol) GetExchangeData(symbol string) map[string]*types.FundingFeeArbitrage {
	wg := new(sync.WaitGroup)
	wg.Add(len(c.Exchanges))

	dataLock := new(sync.Mutex)
	data := map[string]*types.FundingFeeArbitrage{}
	for _, ex := range c.Exchanges {
		go func(ex exchange.Exchange) {
			defer wg.Done()

			info, err := ex.GetFundingAndPrice(symbol)
			if err != nil {
				return
			}
			dataLock.Lock()
			data[ex.Name()] = info
			dataLock.Unlock()
		}(ex)
	}
	wg.Wait()

	return data
}
