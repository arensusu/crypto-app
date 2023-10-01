// https://bitgetlimited.github.io/apidoc/en/mix/
package bitgetEx

import (
	"encoding/json"
	"funding-rate/exchange/strategy"

	"github.com/outtoin/bitget-golang-sdk-api"
)

type BitgetExchange struct {
	Name   string
	Client *bitget.Client
}

func New() *BitgetExchange {
	client := bitget.NewClient()
	return &BitgetExchange{
		Name:   "Bitget",
		Client: client,
	}
}

func (ex *BitgetExchange) GetCrossExArbitrageResponse(coin string) (*strategy.CrossExArbitrageResponse, error) {
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

	return &strategy.CrossExArbitrageResponse{
		ExchangeName:    ex.Name,
		LastPrice:       ticker["data"].(map[string]interface{})["last"].(string),
		FundingRate:     ticker["data"].(map[string]interface{})["fundingRate"].(string),
		NextFundingTime: nextFundingTime["data"].(map[string]interface{})["fundingTime"].(string),
	}, nil
}
