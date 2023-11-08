// https://bitgetlimited.github.io/apidoc/en/mix/
package bitget

import (
	"crypto-exchange/exchange/types"
)

func (ex *Bitget) GetFundingAndPrice(symbol string) (*types.FundingFeeArbitrage, error) {
	bitgetSymbol := symbol + "_UMCBL"
	tickerRes, err := ex.Client.MixMarket().GetTicker(bitgetSymbol)
	if err != nil {
		return nil, err
	}

	fundingTimeRes, err := ex.Client.MixMarket().GetNextFundingTime(bitgetSymbol)
	if err != nil {
		return nil, err
	}

	ticker := tickerRes.Data
	fundingTime := fundingTimeRes.Data

	result := &types.FundingFeeArbitrage{
		LastPrice:   ticker.Last,
		FundingRate: ticker.FundingRate,
		FundingTime: fundingTime.FundingTime,
		Bid1Price:   ticker.BestBid,
		Bid1Size:    ticker.BidSz,
		Ask1Price:   ticker.BestAsk,
		Ask1Size:    ticker.AskSz,
	}
	return result, nil
}
