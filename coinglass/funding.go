package coinglass

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	ApiEndpoint = "https://open-api.coinglass.com/public/v2"
)

var ApiKey = os.Getenv("COINGLASS_APIKEY")

type Pair struct {
	Exchange string
	Symbol   string
}

type FundingRateResponse struct {
	Code    int               `json:"code"`
	Msg     string            `json:"msg"`
	Data    []FundingRateData `json:"data"`
	Success bool              `json:"success"`
}

type FundingRateData struct {
	Rate      float64 `json:"fundingRate"`
	TimeStamp int64   `json:"createTime"`
}

func (pair Pair) GetFundingRate(interval string, return_limit int) (FundingRateResponse, error) {
	url := fmt.Sprintf("%s/indicator/funding?ex=%s&pair=%s&interval=%s&limit=%d", ApiEndpoint, pair.Exchange, pair.Symbol, interval, return_limit)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("coinglassSecret", ApiKey)

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

func (pair Pair) ToMessage(res *FundingRateResponse) string {
	data := res.Data
	length := len(data)
	msg := fmt.Sprintf("\n%s %s\n", pair.Exchange, pair.Symbol)

	if length >= 100 {
		total := totalFundingRate(data, 100)
		msg += fmt.Sprintf("Total of last 100: %.4f%%, APR: %.2f%%\n", total, total/100*3*365)
	}

	if length >= 30 {
		total := totalFundingRate(data, 30)
		msg += fmt.Sprintf("Total of last 30:  %.4f%%, APR: %.2f%%\n", total, total/30*3*365)
	}

	msg += fmt.Sprintf("Last: %.4f%%\n", data[length-1].Rate)

	return msg
}

func totalFundingRate(data []FundingRateData, period int) float64 {
	total := 0.0
	for _, fundingRate := range data[len(data)-period:] {
		total += fundingRate.Rate
	}
	return total
}

func (pair Pair) IsExist() (bool, error) {
	response, err := pair.GetFundingRate("h8", 1)
	if err != nil {
		return false, err
	}
	return response.Msg != "pair unknown", nil
}
