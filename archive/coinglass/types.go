package coinglass

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type CoinglassApi struct {
	Endpoint string
	ApiKey   string
}

const (
	ApiEndpoint = "https://open-api.coinglass.com/public/v2"
)

func NewCoinglassApi(endpoint, apiKey string) CoinglassApi {
	return CoinglassApi{Endpoint: endpoint, ApiKey: apiKey}
}

type APIResponse struct {
	Code    string      `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

func GetDataOfResponse(body []byte, data any) error {
	var response APIResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return err
	}

	if response.Code == "50001" {
		return errors.New("over the maximum of api request")
	}

	dataJSON, err := json.Marshal(response.Data)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(dataJSON, &data); err != nil {
		return err
	}
	return nil
}

func (api CoinglassApi) Request(method, url string) ([]byte, error) {
	req, _ := http.NewRequest(method, url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("coinglassSecret", api.ApiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

type IndicatorFunding struct {
	Rate      float64 `json:"fundingRate"`
	TimeStamp int64   `json:"createTime"`
}

type PerpetualMarket struct {
	ExchangeName   string
	OriginalSymbol string
	Symbol         string
	Price          float64
	Type           int
	UpdateTime     int64
	QuoteCurrency  string
	TurnoverNumber int
	LongRate       float64
	LongVolUsd     float64
	ShortRate      float64
	ShortVolUsd    float64
	// ExchangeLogo string
	// SymbolLogo string
	TotalVolUsd        float64
	Rate               float64
	HighPrice          float64
	LowPrice           float64
	OpenInterestAmount float64
	OpenInterest       float64
	OpenPrice          float64
	PriceChange        float64
	PriceChangePercent float64
	IndexPrice         float64
	BuyTurnoverNumber  float64
	SellTurnoverNumber float64
	FundingRate        float64
	NextFundingTime    int64
}

type FundingRateHistory struct {
	DataList []int                `json:"dataList"`
	DataMap  map[string][]float64 `json:"dataMap"`
}
