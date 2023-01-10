package responses

type EquityPositions struct {
	EquityPositions []EquityPosition `json:"equityPositions"`
}

type EquityPosition struct {
	Symbol                   string `json:"symbol"`
	Name                     string `json:"name"`
	OpenQty                  string `json:"openQty"`
	AvailableForTradingQty   string `json:"availableForTradingQty"`
	AveragePrice             string `json:"averagePrice"`
	MarketValue              string `json:"marketValue"`
	MarketPrice              string `json:"marketPrice"` // i.e. lastTrade price
	PriorClose               string `json:"priorClose"`
	DayProfitOrLoss          string `json:"dayProfitOrLoss"`
	DayProfitOrLossPercent   string `json:"dayProfitOrLossPercent"`
	TotalProfitOrLoss        string `json:"totalProfitOrLoss"`
	TotalProfitOrLossPercent string `json:"totalProfitOrLossPercent"`
}
