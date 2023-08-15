package coinglass

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (api CoinglassApi) GetFundingRate(pair Pair, interval string, return_limit int) (FundingRateResponse, error) {
	url := fmt.Sprintf("%s/indicator/funding?ex=%s&pair=%s&interval=%s&limit=%d", ApiEndpoint, pair.Exchange, pair.Symbol, interval, return_limit)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("coinglassSecret", api.ApiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return FundingRateResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return FundingRateResponse{}, err
	}
	var resBody FundingRateResponse
	json.Unmarshal(body, &resBody)
	return resBody, nil
}
