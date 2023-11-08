package types

type FundingFeeArbitrage struct {
	LastPrice   string `json:"lastPrice"`
	FundingRate string `json:"fundingRate"`
	FundingTime string `json:"fundingTime"`
	Bid1Price   string `json:"bid1Price"`
	Bid1Size    string `json:"bid1Size"`
	Ask1Price   string `json:"ask1Price"`
	Ask1Size    string `json:"ask1Size"`
}
