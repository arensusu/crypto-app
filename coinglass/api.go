package coinglass

import (
	"fmt"
)

func (api CoinglassApi) GetFundingRate(exchange, symbol, interval string, return_limit int) ([]IndicatorFunding, error) {
	url := fmt.Sprintf("%s/indicator/funding?ex=%s&pair=%s&interval=%s&limit=%d", ApiEndpoint, exchange, symbol, interval, return_limit)
	body, err := api.Request("GET", url)
	if err != nil {
		return []IndicatorFunding{}, err
	}

	var data []IndicatorFunding
	if err := GetDataOfResponse(body, &data); err != nil {
		return []IndicatorFunding{}, err
	}
	return data, err
}

func (api CoinglassApi) GetPerpetualMarket(symbol string) ([]PerpetualMarket, error) {
	url := fmt.Sprintf("%s/perpetual_market?symbol=%s", ApiEndpoint, symbol)
	body, err := api.Request("GET", url)
	if err != nil {
		return []PerpetualMarket{}, err
	}

	var data map[string][]PerpetualMarket
	if err := GetDataOfResponse(body, &data); err != nil {
		return []PerpetualMarket{}, err
	}
	return data[symbol], err
}

func (api CoinglassApi) GetFundingRateUSDHistory(symbol, time_period string) (FundingRateHistory, error) {
	url := fmt.Sprintf("%s/funding_usd_history?symbol=%s&time_type=%s", api.Endpoint, symbol, time_period)
	body, err := api.Request("GET", url)
	if err != nil {
		return FundingRateHistory{}, err
	}

	var data FundingRateHistory
	if err := GetDataOfResponse(body, &data); err != nil {
		return FundingRateHistory{}, err
	}
	return data, err
}
