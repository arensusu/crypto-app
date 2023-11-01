package domain

import (
	"strconv"
)

type GetFundingAndPricer interface {
	GetFundingAndPrice(symbol string) (FundingPrice, error)
	GetFundingAndPrices() (FundingPricesOfSymbol, error)
}

type FundingPrice struct {
	Exchange    string
	Price       float64
	FundingRate float64
	FundingTime int64
}

type FundingPricesOfSymbol map[string][]FundingPrice

func (m FundingPricesOfSymbol) Set(exchange, symbol, price, fundingRate, fundingTime string) {
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

func (m FundingPricesOfSymbol) SetSpecial(exchange, symbol, fundingRate, fundingTime string) {
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