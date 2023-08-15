package coinglass

type CoinglassApi struct {
	Endpoint string
	ApiKey   string
}

type Pair struct {
	Exchange string
	Symbol   string
}

const (
	ApiEndpoint = "https://open-api.coinglass.com/public/v2"
)

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

func NewCoinglassApi(endpoint, apiKey string) CoinglassApi {
	return CoinglassApi{Endpoint: endpoint, ApiKey: apiKey}
}
