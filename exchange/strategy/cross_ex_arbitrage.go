package strategy

import (
	"encoding/json"
	"fmt"
)

type CrossExArbitrageResponse struct {
	ExchangeName    string
	LastPrice       string
	FundingRate     string
	NextFundingTime string
}

type CrossExArbitrage interface {
	GetCrossExArbitrageResponse(string) (*CrossExArbitrageResponse, error)
}

func CrossExArbitrageEncoder(exchange string, data interface{}) (*CrossExArbitrageResponse, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return &CrossExArbitrageResponse{}, fmt.Errorf("marshal ticker data failed: %w", err)
	}

	response := new(CrossExArbitrageResponse)
	if err := json.Unmarshal(jsonData, response); err != nil {
		return &CrossExArbitrageResponse{}, fmt.Errorf("unmarshal data to response failed: %w", err)
	}
	response.ExchangeName = exchange
	return response, nil
}
