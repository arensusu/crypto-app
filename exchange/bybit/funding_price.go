// https://bybit-exchange.github.io/docs/v5/intro
package bybit

import (
	"crypto-exchange/exchange/types"
	"errors"

	"github.com/hirokisan/bybit/v2"
)

func (ex *Bybit) GetFundingAndPrice(symbol string) (*types.FundingFeeArbitrage, error) {
	param := bybit.V5GetTickersParam{
		Category: bybit.CategoryV5Linear,
		Symbol:   (*bybit.SymbolV5)(&symbol),
	}
	res, err := ex.Client.V5().Market().GetTickers(param)

	if err != nil {
		return nil, err
	}

	tickers := res.Result.LinearInverse.List
	if len(tickers) != 1 {
		return nil, errors.New("response data length error")
	}

	ticker := tickers[0]

	result := &types.FundingFeeArbitrage{
		LastPrice:   ticker.LastPrice,
		FundingRate: ticker.FundingRate,
		FundingTime: ticker.NextFundingTime,
		Bid1Price:   ticker.Bid1Price,
		Bid1Size:    ticker.Bid1Size,
		Ask1Price:   ticker.Ask1Price,
		Ask1Size:    ticker.Ask1Size,
	}

	return result, nil
}
