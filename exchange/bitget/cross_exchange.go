// https://bitgetlimited.github.io/apidoc/en/mix/
package bitget

import (
	"crypto-exchange/exchange/strategy"
	"encoding/json"
	"strconv"
)

func (ex *Bitget) GetCrossExArbitrageInformation(coin string) (*strategy.CrossExArbitrageInformation, error) {
	symbol := coin + "USDT_UMCBL"
	marketService := ex.Client.GetMixMarketService()

	nextFundingTimeJson, err := marketService.FundingTime(symbol)
	if err != nil {
		return nil, err
	}
	var nextFundingTime map[string]interface{}
	json.Unmarshal([]byte(nextFundingTimeJson), &nextFundingTime)
	if err != nil {
		return nil, err
	}

	tickerJson, err := marketService.Ticker(symbol)
	if err != nil {
		return nil, err
	}
	var ticker map[string]interface{}
	json.Unmarshal([]byte(tickerJson), &ticker)
	if err != nil {
		return nil, err
	}

	price, _ := strconv.ParseFloat(ticker["data"].(map[string]interface{})["last"].(string), 64)
	fundingRate, _ := strconv.ParseFloat(ticker["data"].(map[string]interface{})["fundingRate"].(string), 64)
	fundingTime, _ := strconv.ParseInt(nextFundingTime["data"].(map[string]interface{})["fundingTime"].(string), 10, 64)
	return &strategy.CrossExArbitrageInformation{
		ExchangeName:    ex.Name,
		LastPrice:       price,
		FundingRate:     fundingRate,
		NextFundingTime: fundingTime,
	}, nil
}
